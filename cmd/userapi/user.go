package main

import (
	"log/slog"
	"net/http"

	"github.com/SLANGERES/Tournament-Lederboard/config"
	"github.com/SLANGERES/Tournament-Lederboard/internal/user/handler"
)

func main() {

	cnf := config.SetConfig()

	router := http.NewServeMux()

	//! Routers Endpoints
	router.HandleFunc("POST /v1/user/sign-up", handler.SignInUser())
	router.HandleFunc("POST /v1/user/log-in", handler.LogInUser())

	err := http.ListenAndServe(
		cnf.HttpServer.UserAddress,
		router,
	)
	if err != nil {
		slog.Info("User Server start fail !")
	}
	slog.Info("Server started at " + cnf.HttpServer.UserAddress)

}
