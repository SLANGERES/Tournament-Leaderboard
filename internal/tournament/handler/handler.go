package handler

import (
	"net/http"

	_"github.com/SLANGERES/Tournament-Lederboard/internal/tournament/service"
)

func CreateTournament() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func OngoingTournament() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
