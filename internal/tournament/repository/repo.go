package repository

import (
	"database/sql"
	"errors"

	"github.com/SLANGERES/Tournament-Lederboard/internal/tournament/models"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type TournamentStorage struct {
	Db *sql.DB
}

// ConfigureTournamentStorage initializes DB and tables
func ConfigureTournamentStorage(path string) (*TournamentStorage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// Create tables
	schema := []string{
		`CREATE TABLE IF NOT EXISTS tournaments (
			tournament_id TEXT PRIMARY KEY,
			creator_id TEXT NOT NULL,
			name TEXT NOT NULL,
			description TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS problems (
			problem_id INTEGER PRIMARY KEY AUTOINCREMENT,
			tournament_id TEXT NOT NULL,
			title TEXT NOT NULL,
			description TEXT NOT NULL,
			max_score INTEGER NOT NULL,
			FOREIGN KEY (tournament_id) REFERENCES tournaments(tournament_id) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS test_cases (
			test_case_id INTEGER PRIMARY KEY AUTOINCREMENT,
			problem_id INTEGER NOT NULL,
			input TEXT NOT NULL,
			output TEXT NOT NULL,
			FOREIGN KEY (problem_id) REFERENCES problems(problem_id) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS tournament_participants (
			tournament_id TEXT NOT NULL,
			user_id TEXT NOT NULL,
			user_name TEXT NOT NULL,
			joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			score INTEGER DEFAULT 0,
			PRIMARY KEY (tournament_id, user_id),
			FOREIGN KEY (tournament_id) REFERENCES tournaments(tournament_id) ON DELETE CASCADE
		);`,
	}

	for _, stmt := range schema {
		if _, err := db.Exec(stmt); err != nil {
			return nil, err
		}
	}

	return &TournamentStorage{Db: db}, nil
}

// CreateTournament inserts a new tournament
func (t *TournamentStorage) CreateTournament(creatorID, name, description string) (string, error) {
	tournamentID := uuid.New().String()

	query := `
		INSERT INTO tournaments (tournament_id, creator_id, name, description)
		VALUES (?, ?, ?, ?)
	`

	_, err := t.Db.Exec(query, tournamentID, creatorID, name, description)
	if err != nil {
		return "", err
	}

	return tournamentID, nil
}

// GetAllTournaments returns all tournaments
func (t *TournamentStorage) GetAllTournaments() ([]models.Tournament, error) {
	rows, err := t.Db.Query(`SELECT tournament_id, creator_id, name, description, created_at FROM tournaments`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tournaments []models.Tournament
	for rows.Next() {
		var tr models.Tournament
		if err := rows.Scan(&tr.ID, &tr.CreatorID, &tr.Name, &tr.Description, &tr.CreatedAt); err != nil {
			return nil, err
		}
		tournaments = append(tournaments, tr)
	}

	return tournaments, nil
}

// AddProblem inserts a problem to a tournament
func (t *TournamentStorage) AddProblem(tournamentID string, problem models.Problem) (bool, error) {
	query := `
		INSERT INTO problems (tournament_id, title, description, max_score)
		VALUES (?, ?, ?, ?)
	`
	_, err := t.Db.Exec(query, tournamentID, problem.Title, problem.Description, problem.MaxScore)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetProblems returns all problems for a tournament
func (t *TournamentStorage) GetProblems(tournamentID string) ([]models.GetProblemResponse, error) {
	rows, err := t.Db.Query(`SELECT problem_id, title, description, max_score FROM problems WHERE tournament_id = ?`, tournamentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var problems []models.GetProblemResponse
	for rows.Next() {
		var p models.GetProblemResponse
		if err := rows.Scan(&p.Id, &p.Title, &p.Description, &p.MaxScore); err != nil {
			return nil, err
		}
		problems = append(problems, p)
	}
	return problems, nil
}

// AddTestCase adds a test case to a problem
func (t *TournamentStorage) AddTestCase(problemID string, testCase models.TestCase) (bool, error) {
	query := `
		INSERT INTO test_cases (problem_id, input, output)
		VALUES (?, ?, ?)
	`
	_, err := t.Db.Exec(query, problemID, testCase.Input, testCase.Output)
	if err != nil {
		return false, err
	}
	return true, nil
}

// AddParticipant adds a user to a tournament
func (t *TournamentStorage) AddParticipant(tournamentID, userID, userName string) (bool, error) {
	query := `
		INSERT INTO tournament_participants (tournament_id, user_id, user_name)
		VALUES (?, ?, ?)
	`
	_, err := t.Db.Exec(query, tournamentID, userID, userName)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetAllParticipants returns all participants of a tournament
func (t *TournamentStorage) GetAllParticipants(tournamentID string) ([]models.GetTournamentParticipantResponse, error) {
	rows, err := t.Db.Query(`SELECT user_id, user_name, score FROM tournament_participants WHERE tournament_id = ?`, tournamentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []models.GetTournamentParticipantResponse
	for rows.Next() {
		var p models.GetTournamentParticipantResponse
		if err := rows.Scan(&p.UserID, &p.UserName, &p.Score); err != nil {
			return nil, err
		}
		participants = append(participants, p)
	}
	return participants, nil
}

// GetTournamentById fetches a tournament by ID
func (t *TournamentStorage) GetTournamentById(tournamentID string) (*models.Tournament, error) {
	query := `SELECT tournament_id, creator_id, name, description, created_at FROM tournaments WHERE tournament_id = ?`
	row := t.Db.QueryRow(query, tournamentID)

	var tournament models.Tournament
	if err := row.Scan(&tournament.ID, &tournament.CreatorID, &tournament.Name, &tournament.Description, &tournament.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &tournament, nil
}
