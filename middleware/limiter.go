package middleware

import (
	"AuthService/utils"
	"errors"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

var limitter = rate.NewLimiter(rate.Every(time.Minute), 5)

func RateLimiterMiddelware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limitter.Allow() {
			err := errors.New("Request Limit Exceded")
			utils.WriteJSONErrorResponse(w, http.StatusTooManyRequests, err, "Too many Request")
			return
		}
		next.ServeHTTP(w, r)
	})
}
