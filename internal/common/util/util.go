package util

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Success  bool        `json:"success"`
	Response interface{} `json:"response"`
}
type errorResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error"`
}

func HttpResponse(w http.ResponseWriter, statusCode int, message interface{}) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(response{
		Success:  true,    // or success variable
		Response: message, // or message variable
	})
}
func HttpError(w http.ResponseWriter, statusCode int, errors error) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(errorResponse{
		Success: false,  // or success variable
		Error:   errors, // or message variable
	})
}
