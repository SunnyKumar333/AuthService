package middleware

import (
	"fmt"
	"net/http"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received Request:", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
