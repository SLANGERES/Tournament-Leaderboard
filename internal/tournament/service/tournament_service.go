package service

import "github.com/SLANGERES/Tournament-Lederboard/internal/tournament/models"

type TournamentService struct {
	store TournamentStorage
}

// Constructor
func NewTournamentService(store TournamentStorage) *TournamentService {
	return &TournamentService{store: store}
}

// Methods

func (s *TournamentService) CreateTournament(creatorID, name, description string) (string, error) {
	return s.store.CreateTournament(creatorID, name, description)
}

func (s *TournamentService) GetAllTournaments() ([]models.Tournament, error) {
	return s.store.GetAllTournaments()
}

func (s *TournamentService) GetTournamentById(id string) (models.Tournament, error) {
	return s.store.GetTournamentById(id)
}

func (s *TournamentService) AddProblem(tournamentID string, problem models.Problem) (bool, error) {
	return s.store.AddProblem(tournamentID, problem)
}

func (s *TournamentService) GetProblems(tournamentID string) ([]models.GetProblemResponse, error) {
	return s.store.GetProblems(tournamentID)
}

func (s *TournamentService) AddTestCase(problemID string, testCase models.TestCase) (bool, error) {
	return s.store.AddTestCase(problemID, testCase)
}

func (s *TournamentService) AddParticipant(tournamentID, userID string) (bool, error) {
	return s.store.AddParticipant(tournamentID, userID)
}

func (s *TournamentService) GetAllParticipants(tournamentID string) ([]models.GetTournamentParticipantResponse, error) {
	return s.store.GetAllParticipants(tournamentID)
}
