package server

import (
	"encoding/json"
	"fmt"
	"github.com/alexey-zayats/claim-handler/internal/application"
	"github.com/alexey-zayats/claim-handler/internal/form"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"io/ioutil"
	"net/http"
	"time"
)

// ServeSingle ...
func (s *Server) ServeSingle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		msg := fmt.Sprintf(format, "Bad request", "POST method required", 1, 500)
		s.http500Error([]byte(msg), w, r)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{"reason": err}).Error("unable read body")
		msg := fmt.Sprintf(format, "Bad request", "Unable read HTTP body", 2, 500)
		s.http500Error([]byte(msg), w, r)
		return
	}

	form := &form.Single{}

	err = json.Unmarshal(body, form)
	if err != nil {
		logrus.WithFields(logrus.Fields{"reason": err}).Error("unable json.Unmarshal form post")
		e := fmt.Sprintf("Невозможно разобрать данные: %s", err)
		msg := fmt.Sprintf(format, "Bad request", e, 3, 500)
		s.http500Error([]byte(msg), w, r)
		return
	}

	err = s.validate.Struct(form)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		if len(validationErrors) > 0 {
			s.formErrors(validationErrors, w, r)
			return
		}
	}

	app := application.NewSingle(form)
	err = app.Validate()
	if err != nil {
		appErrors := err.(application.ValidationErrors)

		if len(appErrors) > 0 {
			s.appErrors(appErrors, w, r)
			return
		}
	}

	expire := time.Duration(s.conf.Cache.Expire.Single) * time.Minute

	key := fmt.Sprintf("%d-%d", app.Inn, app.Ogrn)
	if _, ok := s.cache.Get(key); ok {
		logrus.WithFields(logrus.Fields{"INN": app.Inn}).Error("rate limit")

		txt := fmt.Sprintf("Вы не можете подавать заявку чаще чем один раз в течение %s", expire.String())
		msg := fmt.Sprintf(format, "Bad request", txt, 100, 400)
		s.http500Error([]byte(msg), w, r)
		return
	}

	s.cache.Set(key, true, expire)

	err = s.que.Publish(s.conf.Amqp.Exchange, s.conf.Amqp.Routing.Single, app, amqp.Table{}, amqp.Table{})
	if err != nil {
		logrus.WithFields(logrus.Fields{"reason": err}).Error("unable send data to queue")
		msg := fmt.Sprintf(format, "Internal server error", "Unable send data to queue", 4, 500)
		s.http500Error([]byte(msg), w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
