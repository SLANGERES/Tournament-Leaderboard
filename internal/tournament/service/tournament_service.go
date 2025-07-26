package service

import (
	"github.com/SLANGERES/Tournament-Lederboard/internal/tournament/models"
	"github.com/SLANGERES/Tournament-Lederboard/internal/tournament/repository"
)

type TournamentService struct {
	store *repository.TournamentStorage
}

func NewTournamentService(store *repository.TournamentStorage) *TournamentService {
	return &TournamentService{store: store}
}

// Service Methods

func (s *TournamentService) CreateTournament(creatorID, name, description string) (string, error) {
	return s.store.CreateTournament(creatorID, name, description)
}

func (s *TournamentService) GetAllTournaments() ([]models.Tournament, error) {
	return s.store.GetAllTournaments()
}

func (s *TournamentService) GetTournamentById(id string) (*models.Tournament, error) {
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

func (s *TournamentService) AddParticipant(tournamentID, userID, userName string) (bool, error) {
	return s.store.AddParticipant(tournamentID, userID, userName)
}

func (s *TournamentService) GetAllParticipants(tournamentID string) ([]models.GetTournamentParticipantResponse, error) {
	return s.store.GetAllParticipants(tournamentID)
}
