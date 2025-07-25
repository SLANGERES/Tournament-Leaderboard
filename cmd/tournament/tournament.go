package main

import (
	"log/slog"
	"net/http"

	"github.com/SLANGERES/Tournament-Lederboard/config"
	"github.com/SLANGERES/Tournament-Lederboard/internal/tournament/handler"
	"github.com/SLANGERES/Tournament-Lederboard/internal/tournament/repository"
	"github.com/SLANGERES/Tournament-Lederboard/internal/tournament/service"
)

func main() {
	cnf, err := config.SetConfig()
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		return
	}

	db, err := repository.ConfigureTournamentStorage(cnf.AdminDB)
	if err != nil {
		slog.Error("Failed to configure DB", "error", err)
		return
	}

	tournamentService := service.NewTournamentService(db)
	router := http.NewServeMux()

	// Corrected endpoint path (typo)
	router.HandleFunc("POST /v1/tournament/", handler.CreateTournament(*tournamentService))
	router.HandleFunc("GET /v1/tournament/ongoing", handler.GetOngoingTournament(*tournamentService))

	slog.Info("Server starting at " + cnf.HttpServer.TournamentAddress)
	err = http.ListenAndServe(cnf.HttpServer.TournamentAddress, router)
	if err != nil {
		slog.Error("Server failed to start", "error", err)
	}
}
