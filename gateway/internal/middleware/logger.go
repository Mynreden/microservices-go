package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Pass the request to the next handler
		next.ServeHTTP(w, r)

		// Log the request details
		log.Printf(
			"%s %s %s %s %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
			r.UserAgent(),
		)
	})
}
