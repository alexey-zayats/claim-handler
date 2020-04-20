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
)

var format = `{"name":"%s","message":"%s","code":%d,"status":%d}`

// ServePeople ...
func (s *Server) ServePeople(w http.ResponseWriter, r *http.Request) {

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

	form := &form.People{}

	err = json.Unmarshal(body, form)
	if err != nil {
		logrus.WithFields(logrus.Fields{"reason": err}).Error("unable json.Unmarshal form post")
		msg := fmt.Sprintf(format, "Bad request", "Unable unmarshal body to json", 3, 500)
		s.http500Error([]byte(msg), w, r)
		return
	}

	err = s.validate.Struct(form)
	validationErrors := err.(validator.ValidationErrors)

	if len(validationErrors) > 0 {
		s.displayErrors(validationErrors, w, r)
		return
	}

	app := application.People(form)

	err = s.que.Publish(s.conf.Amqp.Exchange, s.conf.Amqp.Routing.People, app, amqp.Table{}, amqp.Table{})
	if err != nil {
		logrus.WithFields(logrus.Fields{"reason": err}).Error("unable send data to queue")
		msg := fmt.Sprintf(format, "Internal server error", "Unable send data to queue", 4, 500)
		s.http500Error([]byte(msg), w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
