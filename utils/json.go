package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	Validator = NewValidator()
}

func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	jsonEncoder := json.NewEncoder(w)

	return jsonEncoder.Encode(payload)

}

func WriteJSONSuccessResponse(w http.ResponseWriter, statusCode int, payload any, message string) error {
	response := map[string]any{
		"success": true,
		"data":    payload,
		"error":   nil,
		"message": message,
	}

	return WriteJSONResponse(w, statusCode, response)
}
func WriteJSONErrorResponse(w http.ResponseWriter, statusCode int, err error, message string) error {
	response := map[string]any{
		"success": false,
		"data":    nil,
		"error":   err.Error(),
		"message": message,
	}

	return WriteJSONResponse(w, statusCode, response)
}

func ReadJSONBody(r *http.Request, payload any) error {
	jsonDecoder := json.NewDecoder(r.Body)
	jsonDecoder.DisallowUnknownFields()
	return jsonDecoder.Decode(payload)
}
