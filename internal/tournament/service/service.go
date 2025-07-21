package service

import "github.com/SLANGERES/Tournament-Lederboard/internal/tournament/models"

type TournamentStorage interface {

	// Tournament
	CreateTournament(creatorID, name, description string) (string, error) // return UUID or string ID
	GetAllTournaments() ([]models.Tournament, error)

	// Problems
	AddProblem(tournamentID string, problem models.Problem) (bool, error)
	GetProblems(tournamentID string) ([]models.GetProblemResponse, error)

	// Test Cases
	AddTestCase(problemID string, testCase models.TestCase) (bool, error)

	// Participants
	AddParticipant(tournamentID, userID string) (bool, error)
	GetAllParticipants(tournamentID string) ([]models.GetProblemResponse, error)
}
