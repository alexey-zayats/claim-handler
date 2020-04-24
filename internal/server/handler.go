package server

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func (s *Server) makeHandler() *http.ServeMux {
	handler := http.NewServeMux()

	handler.Handle("/metrics", promhttp.Handler())
	handler.HandleFunc("/handler/vehicle", s.ServeVehicle)
	handler.HandleFunc("/handler/people", s.ServePeople)
	handler.HandleFunc("/handler/single", s.ServeSingle)

	return handler
}
