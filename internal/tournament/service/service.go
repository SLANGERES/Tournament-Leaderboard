package service

import "github.com/SLANGERES/Tournament-Lederboard/internal/tournament/models"

type TournamentStorage interface {

	// Tournament
	CreateTournament(creatorID, name, description string) (string, error)
	GetAllTournaments() ([]models.Tournament, error)
	GetTournamentById(id string) (*models.Tournament, error) // ✅ pointer

	// Problems
	AddProblem(tournamentID string, problem models.Problem) (bool, error)
	GetProblems(tournamentID string) ([]models.GetProblemResponse, error)

	// Test Cases
	AddTestCase(problemID int, testCase models.TestCase) (bool, error) // ✅ int

	// Participants
	AddParticipant(tournamentID, userID, userName string) (bool, error) // ✅ userName added
	GetAllParticipants(tournamentID string) ([]models.GetTournamentParticipantResponse, error)
}
