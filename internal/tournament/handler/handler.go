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

// CreateTournament godoc
// @Summary Create a new tournament
// @Description Admin-only endpoint to create a new tournament
// @Tags tournaments
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT"
// @Param tournament body models.Tournament true "Tournament details"
// @Success 201 {object} models.Tournament
// @Failure 400 {object} util.ErrorResponse
// @Failure 401 {object} util.ErrorResponse
// @Failure 500 {object} util.ErrorResponse
// @Router /tournaments [post]
func CreateTournament(tournamentService service.TournamentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claim, ok := middleware.ClaimsJWT(r.Context())
		if !ok {
			util.HttpError(w, http.StatusUnauthorized, fmt.Errorf("unable to get JWT claims"))
			return
		}
		if claim.Role != "admin" {
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

// GetOngoingTournament godoc
// @Summary Get all ongoing tournaments
// @Description Returns a list of all ongoing tournaments
// @Tags tournaments
// @Produce json
// @Success 200 {array} models.Tournament
// @Failure 500 {object} util.ErrorResponse
// @Router /tournaments/ongoing [get]
func GetOngoingTournament(tournamentService service.TournamentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tournaments, err := tournamentService.GetAllTournaments()
		if err != nil {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("failed to fetch tournaments: %v", err))
			return
		}
		util.HttpResponse(w, http.StatusOK, tournaments)
	}
}

// GetTournamentByID godoc
// @Summary Get tournament by ID
// @Description Returns details of a tournament by its ID
// @Tags tournaments
// @Produce json
// @Param id path string true "Tournament ID"
// @Success 200 {object} models.Tournament
// @Failure 400 {object} util.ErrorResponse
// @Failure 404 {object} util.ErrorResponse
// @Router /tournaments/{id} [get]
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

// AddProblemInTournament godoc
// @Summary Add a problem to a tournament
// @Description Admin-only endpoint to add a coding problem to a tournament
// @Tags tournaments
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT"
// @Param id path string true "Tournament ID"
// @Param problem body models.Problem true "Problem details"
// @Success 200 {string} string "Problem added successfully"
// @Failure 400 {object} util.ErrorResponse
// @Failure 401 {object} util.ErrorResponse
// @Failure 500 {object} util.ErrorResponse
// @Router /tournaments/addproblem/{id} [post]
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

// GetProblem godoc
// @Summary Get problems for a tournament
// @Description Returns a list of problems for a given tournament
// @Tags tournaments
// @Produce json
// @Param id path string true "Tournament ID"
// @Success 200 {array} models.GetProblemResponse
// @Failure 400 {object} util.ErrorResponse
// @Failure 500 {object} util.ErrorResponse
// @Router /tournaments/ongoing/{id}/problem [get]

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

// AddParticipant godoc
// @Summary Add participant to an ongoing tournament
// @Description Allows a logged-in user to join an ongoing tournament
// @Tags tournaments
// @Produce json
// @Param Authorization header string true "Bearer JWT"
// @Param id path string true "Tournament ID"
// @Success 200 {string} string "Participant added successfully"
// @Failure 400 {object} util.ErrorResponse
// @Failure 401 {object} util.ErrorResponse
// @Failure 500 {object} util.ErrorResponse
// @Router /tournaments/ongoing/addparticipant/{id} [post]
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

		ok, err := tournamentService.AddParticipant(tournamentID, claim.Subject, claim.Username)
		if err != nil || !ok {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("could not add participant: %v", err))
			return
		}

		util.HttpResponse(w, http.StatusOK, "Participant added successfully")
	}
}

// GetAllParticipant godoc
// @Summary Get all participants for a tournament
// @Description Returns a list of participants in a given tournament
// @Tags tournaments
// @Produce json
// @Param id path string true "Tournament ID"
// @Success 200 {array} models.GetTournamentParticipantResponse
// @Failure 400 {object} util.ErrorResponse
// @Failure 500 {object} util.ErrorResponse
// @Router /tournaments/{id}/participants [get]
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

// AddNewTestCase godoc
// @Summary Add a new test case to a problem
// @Description Admin-only endpoint to add a test case to a problem
// @Tags tournaments
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT"
// @Param id path string true "Problem ID"
// @Param testcase body models.TestCase true "Test case details"
// @Success 200 {string} string "Test case added successfully"
// @Failure 400 {object} util.ErrorResponse
// @Failure 401 {object} util.ErrorResponse
// @Failure 500 {object} util.ErrorResponse
// @Router /tournaments/addtestcase/{id} [post]
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
