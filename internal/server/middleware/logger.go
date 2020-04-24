package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

// AddLogger logs reqid/response pair
func AddLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("X-Liveness-Probe") == "Healthz" {
			h.ServeHTTP(w, r)
			return
		}

		t1 := time.Now()

		h.ServeHTTP(w, r)

		// Prepare fields to log
		var scheme string
		if r.TLS != nil {
			scheme = "https"
		} else {
			scheme = "http"
		}

		uri := strings.Join([]string{scheme, "://", r.Host, r.RequestURI}, "")

		status := r.Header.Get("Status")
		forwarded := r.Header.Get("X-Forwarded-For")
		method := r.Method
		remoteAddr := r.RemoteAddr
		userAgent := r.UserAgent()

		// Log HTTP response
		logrus.WithFields(logrus.Fields{
			"http-method": method,
			"remote-addr": remoteAddr,
			"forwarded":   forwarded,
			"user-agent":  userAgent,
			"uri":         uri,
			"latency":     time.Since(t1),
			"status":      status,
		}).Info("Request")

	})
}
