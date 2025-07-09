package main

//Admin Backend API

import (
	"log/slog"
	"net/http"

	"github.com/SLANGERES/Tournament-Lederboard/internal/admin/handler"
)

func main() {
	router := http.NewServeMux()

	//! Routers Endpoints
	router.HandleFunc("POST /sign-up", handler.Signup())
	router.HandleFunc("POST /log-in", handler.Login())

	err := http.ListenAndServe(
		"0.0.0.0:7070",
		router,
	)
	if err != nil {
		slog.Info("Server start fail !")
	}
	slog.Info("Admin Server started at 0.0.0.0:7070")

}
