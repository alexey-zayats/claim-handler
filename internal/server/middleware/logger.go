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

		// We do not want to be spammed by Kubernetes health check.
		// Do not log Kubernetes health check.
		// You can change this behavior as you wish.
		if r.Header.Get("X-Liveness-Probe") == "Healthz" {
			h.ServeHTTP(w, r)
			return
		}

		proto := r.Proto
		method := r.Method
		remoteAddr := r.RemoteAddr
		userAgent := r.UserAgent()

		// Prepare fields to log
		var scheme string
		if r.TLS != nil {
			scheme = "https"
		} else {
			scheme = "http"
		}

		uri := strings.Join([]string{scheme, "://", r.Host, r.RequestURI}, "")

		t1 := time.Now()

		h.ServeHTTP(w, r)

		// Log HTTP response
		logrus.WithFields(logrus.Fields{
			"http-scheme": scheme,
			"http-proto":  proto,
			"http-method": method,
			"remote-addr": remoteAddr,
			"user-agent":  userAgent,
			"uri":         uri,
			"latency":     time.Since(t1),
		}).Info("Request")

	})
}
