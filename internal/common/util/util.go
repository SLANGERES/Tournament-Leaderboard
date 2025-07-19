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
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func HttpResponse(w http.ResponseWriter, statusCode int, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(response{
		Success:  true,
		Response: message,
	})
}

func HttpError(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(errorResponse{
		Success: false,
		Error:   err.Error(),
	})
}
