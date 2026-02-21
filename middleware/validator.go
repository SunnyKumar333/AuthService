package middleware

import (
	"AuthService/dto"
	"AuthService/utils"
	"context"
	"net/http"
)

func LoginRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.LoginUserDTO
		if parseError := utils.ReadJSONBody(r, &payload); parseError != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, parseError, "Unable to Parse Request Body")
			return
		}

		if validationError := utils.Validator.Struct(payload); validationError != nil {
			utils.WriteJSONErrorResponse(w, http.StatusBadRequest, validationError, "Invalid Request Body")
			return
		}

		ctx := context.WithValue(r.Context(), "payload", payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
