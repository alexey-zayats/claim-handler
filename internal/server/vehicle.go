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

// ServeVehicle ...
func (s *Server) ServeVehicle(w http.ResponseWriter, r *http.Request) {
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

	form := &form.Vehicle{}

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

	app := application.Vehicle(form)

	err = s.que.Publish(s.conf.Amqp.Exchange, s.conf.Amqp.Routing.Vehicle, app, amqp.Table{}, amqp.Table{})
	if err != nil {
		logrus.WithFields(logrus.Fields{"reason": err}).Error("unable send data to queue")
		msg := fmt.Sprintf(format, "Internal server error", "Unable send data to queue", 4, 500)
		s.http500Error([]byte(msg), w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}

func (s *Server) displayErrors(e validator.ValidationErrors, w http.ResponseWriter, r *http.Request) {

	fields := make(map[string][]string)
	for _, v := range e {
		fields[v.Tag()] = append(fields[v.Tag()], fmt.Sprintf("Ошибка валидации поля '%s'", v.Tag()))
	}

	jsf, err := json.Marshal(fields)
	if err != nil {
		logrus.WithFields(logrus.Fields{"reason": err}).Error("unable marshal fields")
		msg := fmt.Sprintf(format, "Internal server error", "unable marshal fields", 5, 500)
		s.http500Error([]byte(msg), w, r)
		return
	}

	msg := fmt.Sprintf(`{"errors": %s}`, jsf)
	s.http500Error([]byte(msg), w, r)
}
