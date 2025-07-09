package handler

import "net/http"

func CreateTournament() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func OngoingTournament() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
