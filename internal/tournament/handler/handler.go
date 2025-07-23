package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SLANGERES/Tournament-Lederboard/internal/common/middleware"
	"github.com/SLANGERES/Tournament-Lederboard/internal/common/util"
	"github.com/SLANGERES/Tournament-Lederboard/internal/tournament/models"
	"github.com/SLANGERES/Tournament-Lederboard/internal/tournament/repository"
	_ "github.com/SLANGERES/Tournament-Lederboard/internal/tournament/service"
	"github.com/go-playground/validator/v10"
)

func CreateTournament(tounamentStorage repository.TournamentStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claim, ok := middleware.ClaimsJWT(r.Context())
		if !ok {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("unable to get the claims"))
			return
		}
		var newTournament models.Tournament
		err := json.NewDecoder(r.Body).Decode(&newTournament)
		if err != nil {
			//error occure
		}
		err = validator.New().Struct(newTournament)
		if err != nil {
			//validation error
		}
		data, err := tounamentStorage.CreateTournament(claim.ID, newTournament.Name, newTournament.Description)

		if err != nil {
			//DB error
		}
		util.HttpResponse(w, http.StatusCreated, data)
	}
}
func GetOngoingTournament(tounamentStorage repository.TournamentStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tournaments []models.Tournament
		var err error
		tournaments, err = tounamentStorage.GetTournamentById(id)

		if err != nil {
			//db error
		}
		util.HttpResponse(w, http.StatusOK, tournaments)
	}
}
func GetTournamentByid(db repository.TournamentStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("tournamentID")

		var tournament models.Tournament
		var err error
		tournament, err = db.GetTournamentBsyId(id)

	}
}
