package main


//Tournament API Backend

import (
	"log/slog"
	"net/http"

	"github.com/SLANGERES/Tournament-Lederboard/internal/tournament/handler"
)

func main() {
	router := http.NewServeMux()

	//! Routers Endpoints
	router.HandleFunc("POST /v1/tourament/", handler.CreateTournament())
	router.HandleFunc("POST /v1/tournament/", handler.OngoingTournament())

	err := http.ListenAndServe(
		"0.0.0.0:7071",
		router,
	)
	if err != nil {
		slog.Info("Server start fail !")
	}
	slog.Info("Server started at 0.0.0.0:7071")

}
