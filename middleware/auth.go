package middleware

import (
	env "AuthService/config/env"
	"AuthService/utils"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
)

func JWTAuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//check if Authorization Header Exist
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			err := errors.New("Authorization Header Not Found")
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, err, "Token Not Found")
			return
		}

		jwtToken, hasBearerPrefix := strings.CutPrefix(authHeader, "Bearer ")

		if !hasBearerPrefix {
			err := errors.New("Invalid Token Formet")
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, err, "Bearer Token Not Found")
			return
		}
		claims := jwt.MapClaims{}
		_, jwtValidationError := jwt.ParseWithClaims(jwtToken, &claims, func(t *jwt.Token) (any, error) {
			return []byte(env.GetString("JWT_SECRET", "top_secret")), nil
		})

		if jwtValidationError != nil {
			utils.WriteJSONErrorResponse(w, http.StatusUnauthorized, jwtValidationError, "Invalid token")
			return
		}
		fmt.Println("My claims:", claims)
		idFloat, _ := claims["id"].(float64)
		userEmail, _ := claims["email"]

		userId := strconv.Itoa(int(idFloat))

		fmt.Println("id:", userId, "Email:", userEmail)
		ctx := context.WithValue(r.Context(), "userId", userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
