package main

import (
	"log/slog"
	"net/http"

	"github.com/SLANGERES/Tournament-Lederboard/internal/user/handler"
)

func main() {
	router := http.NewServeMux()

	//! Routers Endpoints
	router.HandleFunc("POST /v1/user/sign-up", handler.SignInUser())
	router.HandleFunc("POST /v1/user/log-in", handler.LogInUser())

	err := http.ListenAndServe(
		"0.0.0.0:7072",
		router,
	)
	if err != nil {
		slog.Info("User Server start fail !")
	}
	slog.Info("Server started at 0.0.0.0:7072")

}
