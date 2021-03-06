package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alexey-zayats/claim-handler/internal/application"
	"github.com/alexey-zayats/claim-handler/internal/config"
	"github.com/alexey-zayats/claim-handler/internal/queue"
	"github.com/alexey-zayats/claim-handler/internal/server/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"net"
	"net/http"
	"strings"
	"time"
)

// Server структура данных сервера
type Server struct {
	conf    *config.Config
	que     *queue.Queue
	httpSrv *http.Server
	serving bool

	validate *validator.Validate
	cache    *cache.Cache
	expire   time.Duration
}

// DI структура параметров сервера
type DI struct {
	dig.In
	Config *config.Config
	Queue  *queue.Queue
}

var format = `{"name":"%s","message":"%s","code":%d,"status":%d}`

// NewServer метод конструктора сервера
func NewServer(di DI) *Server {

	expire := time.Duration(di.Config.Cache.Expire.Default) * time.Minute
	cleanup := time.Duration(di.Config.Cache.Cleanup) * time.Minute

	validate := validator.New()

	s := &Server{
		conf:     di.Config,
		que:      di.Queue,
		validate: validate,
		cache:    cache.New(expire, cleanup),
	}
	return s
}

// Healthy ...
func (s *Server) Healthy() (bool, error) {
	if false == s.serving {
		return false, fmt.Errorf("not serving")
	}
	return s.serving, nil
}

// Start метод запуска сервера
func (s *Server) Start(ctx context.Context) error {

	listen := fmt.Sprintf("%s:%d", s.conf.Listen.Host, s.conf.Listen.Port)

	s.httpSrv = &http.Server{
		Addr: listen,
		// add handler with middleware
		Handler:           middleware.AddLogger(s.makeHandler()),
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       5 * time.Second,
	}

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			logrus.WithFields(logrus.Fields{"listen": listen}).Info("shutting down the server...")
			if err := s.httpSrv.Shutdown(ctx); err != nil {
				logrus.WithFields(logrus.Fields{"reason": err}).Error("unable shutdown")
			}
			s.serving = false
		}
	}(ctx)

	listener, err := net.Listen("tcp", listen)
	if err != nil {
		return errors.Wrapf(err, "unable listen %s", listen)
	}

	s.serving = true

	logrus.WithFields(logrus.Fields{"listen": listen}).Info("starting server...")

	return s.httpSrv.Serve(listener)
}

func (s *Server) http500Error(err []byte, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(err)
}

func (s *Server) formErrors(e validator.ValidationErrors, w http.ResponseWriter, r *http.Request) {

	fields := make(map[string][]string)
	for _, v := range e {

		logrus.WithFields(logrus.Fields{
			"Kind":            v.Kind(),
			"Value":           v.Value(),
			"Tag":             v.Tag(),
			"StructField":     v.StructField(),
			"Field":           v.Field(),
			"Type":            v.Type(),
			"ActualTag":       v.ActualTag(),
			"Namespace":       v.Namespace(),
			"Param":           v.Param(),
			"StructNamespace": v.StructNamespace(),
		}).Error("form.validation")

		f := strings.ToLower(v.Field())
		fields[f] = append(fields[v.Tag()], fmt.Sprintf("Ошибка валидации поля '%s': %s %s", f, v.Tag(), v.Param()))
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

func (s *Server) appErrors(ve application.ValidationErrors, w http.ResponseWriter, r *http.Request) {

	for k, v := range ve {
		logrus.WithFields(logrus.Fields{
			k: v,
		}).Error("app.validation")
	}

	jsf, err := json.Marshal(ve)
	if err != nil {
		logrus.WithFields(logrus.Fields{"reason": err}).Error("unable marshal fields")
		msg := fmt.Sprintf(format, "Internal server error", "unable marshal fields", 6, 500)
		s.http500Error([]byte(msg), w, r)
		return
	}

	msg := fmt.Sprintf(`{"errors": %s}`, jsf)
	s.http500Error([]byte(msg), w, r)
}
