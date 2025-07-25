package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SLANGERES/Tournament-Lederboard/internal/common/middleware"
	"github.com/SLANGERES/Tournament-Lederboard/internal/common/util"
	"github.com/SLANGERES/Tournament-Lederboard/internal/tournament/models"
	"github.com/SLANGERES/Tournament-Lederboard/internal/tournament/service"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateTournament(tournamentService service.TournamentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claim, ok := middleware.ClaimsJWT(r.Context())
		if !ok {
			util.HttpError(w, http.StatusUnauthorized, fmt.Errorf("unable to get JWT claims"))
			return
		}

		var newTournament models.Tournament
		if err := json.NewDecoder(r.Body).Decode(&newTournament); err != nil {
			util.HttpError(w, http.StatusBadRequest, fmt.Errorf("invalid request body"))
			return
		}

		if err := validate.Struct(newTournament); err != nil {
			util.HttpError(w, http.StatusBadRequest, fmt.Errorf("validation error: %v", err))
			return
		}

		data, err := tournamentService.CreateTournament(claim.ID, newTournament.Name, newTournament.Description)
		if err != nil {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("could not create tournament: %v", err))
			return
		}

		util.HttpResponse(w, http.StatusCreated, data)
	}
}

func GetOngoingTournament(tournamentService *service.TournamentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tournaments, err := tournamentService.GetAllTournaments()
		if err != nil {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("failed to fetch tournaments: %v", err))
			return
		}
		util.HttpResponse(w, http.StatusOK, tournaments)
	}
}

func GetTournamentByID(tournamentService service.TournamentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("tournamentID")
		if id == "" {
			util.HttpError(w, http.StatusBadRequest, fmt.Errorf("tournament ID is required"))
			return
		}

		tournament, err := tournamentService.GetTournamentById(id)
		if err != nil {
			util.HttpError(w, http.StatusNotFound, fmt.Errorf("tournament not found"))
			return
		}

		util.HttpResponse(w, http.StatusOK, tournament)
	}
}

func AddProblemInTournament(tournamentService service.TournamentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			util.HttpError(w, http.StatusBadRequest, fmt.Errorf("tournament ID is required"))
			return
		}

		var problem models.Problem
		if err := json.NewDecoder(r.Body).Decode(&problem); err != nil {
			util.HttpError(w, http.StatusBadRequest, fmt.Errorf("invalid request body"))
			return
		}

		if err := validate.Struct(problem); err != nil {
			util.HttpError(w, http.StatusBadRequest, fmt.Errorf("validation error: %v", err))
			return
		}

		ok, err := tournamentService.AddProblem(id, problem)
		if err != nil || !ok {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("failed to add problem: %v", err))
			return
		}

		util.HttpResponse(w, http.StatusOK, "Problem added successfully")
	}
}

func GetProblem(tournamentService service.TournamentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			util.HttpError(w, http.StatusBadRequest, fmt.Errorf("tournament ID is required"))
			return
		}

		problems, err := tournamentService.GetProblems(id)
		if err != nil {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("failed to fetch problems: %v", err))
			return
		}

		util.HttpResponse(w, http.StatusOK, problems)
	}
}

func AddParticipant(tournamentService service.TournamentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claim, ok := middleware.ClaimsJWT(r.Context())
		if !ok {
			util.HttpError(w, http.StatusUnauthorized, fmt.Errorf("unable to get JWT claims"))
			return
		}

		tournamentID := r.PathValue("id")
		if tournamentID == "" {
			util.HttpError(w, http.StatusBadRequest, fmt.Errorf("tournament ID is required"))
			return
		}

		ok, err := tournamentService.AddParticipant(tournamentID, claim.Subject)
		if err != nil || !ok {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("could not add participant: %v", err))
			return
		}

		util.HttpResponse(w, http.StatusOK, "Participant added successfully")
	}
}

func GetAllParticipant(tournamentService service.TournamentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			util.HttpError(w, http.StatusBadRequest, fmt.Errorf("tournament ID is required"))
			return
		}

		participant, err := tournamentService.GetAllParticipants(id)
		if err != nil {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("failed to fetch participants: %v", err))
			return
		}

		util.HttpResponse(w, http.StatusOK, participant)
	}
}

func AddNewTestCase(tournamentService service.TournamentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			util.HttpError(w, http.StatusBadRequest, fmt.Errorf("problem ID is required"))
			return
		}

		var newTestCase models.TestCase
		if err := json.NewDecoder(r.Body).Decode(&newTestCase); err != nil {
			util.HttpError(w, http.StatusBadRequest, fmt.Errorf("invalid request body"))
			return
		}

		if err := validate.Struct(newTestCase); err != nil {
			util.HttpError(w, http.StatusBadRequest, fmt.Errorf("validation error: %v", err))
			return
		}

		ok, err := tournamentService.AddTestCase(id, newTestCase)
		if err != nil || !ok {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("failed to add test case: %v", err))
			return
		}

		util.HttpResponse(w, http.StatusOK, "Test case added successfully")
	}
}
