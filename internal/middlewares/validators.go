package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/ayushWeb07/AirBnb-Go-Api-Gateway/internal/utils"
)

// HTTP middleware to decode JSON data
func DecodeRequestBody[T any](next http.Handler) http.Handler {
	return http.HandlerFunc(func(resWriter http.ResponseWriter, req *http.Request) {
		userPayload := new(T)

		// read the data from the request body
		decodeErr := json.NewDecoder(req.Body).Decode(&userPayload)

		if decodeErr != nil {
			utils.WriteJsonResponse(http.StatusBadRequest, resWriter, map[string]any{
				"success": false,
				"message": "Failed to decode the json body",
				"error":   decodeErr.Error(),
			})

			return
		}

		next.ServeHTTP(resWriter, req)
	})
}
