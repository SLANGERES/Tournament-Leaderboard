package main

//Tournament API Backend

import (
	
	"log/slog"
	"net/http"

	"github.com/SLANGERES/Tournament-Lederboard/config"
	"github.com/SLANGERES/Tournament-Lederboard/internal/tournament/handler"
	"github.com/SLANGERES/Tournament-Lederboard/internal/tournament/repository"
)

func main() {
	cnf, err := config.SetConfig()
	router := http.NewServeMux()

	db, err := repository.ConfigureTournamentStorage(cnf.AdminDB)

	//! Routers Endpoints
	router.HandleFunc("POST /v1/tourament/", handler.CreateTournament())
	
	router.HandleFunc("POST /v1/tournament/", handler.OngoingTournament())

	err = http.ListenAndServe(
		cnf.HttpServer.TournamentAddress,
		router,
	)
	if err != nil {
		slog.Info("Server start fail !")
	}
	slog.Info("Server started at" + cnf.HttpServer.TournamentAddress)

}
