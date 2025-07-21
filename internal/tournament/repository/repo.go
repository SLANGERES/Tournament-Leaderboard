package repository

import (
	"database/sql"

	"github.com/SLANGERES/Tournament-Lederboard/internal/tournament/models"
	_ "github.com/mattn/go-sqlite3"
)

type TournamentStorage struct {
	Db *sql.DB
}

func ConfigureTournamentStorage(path string) (*TournamentStorage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// Create tournaments table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tournaments (
		tournament_id TEXT PRIMARY KEY,
		creator_id TEXT NOT NULL,
		name TEXT NOT NULL,
		description TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		return nil, err
	}

	// Create problems table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS problems (
		problem_id INTEGER PRIMARY KEY AUTOINCREMENT,
		tournament_id TEXT NOT NULL,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		max_score INTEGER NOT NULL,
		FOREIGN KEY (tournament_id) REFERENCES tournaments(tournament_id) ON DELETE CASCADE
	);`)
	if err != nil {
		return nil, err
	}

	// Create test_cases table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS test_cases (
		test_case_id INTEGER PRIMARY KEY AUTOINCREMENT,
		problem_id INTEGER NOT NULL,
		input TEXT NOT NULL,
		output TEXT NOT NULL,
		FOREIGN KEY (problem_id) REFERENCES problems(problem_id) ON DELETE CASCADE
	);`)
	if err != nil {
		return nil, err
	}

	// Create tournament_participants table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tournament_participants (
		tournament_id TEXT NOT NULL,
		user_id TEXT NOT NULL,
		user_name TEXT NOT NULL,
		joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		score INTEGER DEFAULT 0,
		PRIMARY KEY (tournament_id, user_id),
		FOREIGN KEY (tournament_id) REFERENCES tournaments(tournament_id) ON DELETE CASCADE
	);`)
	if err != nil {
		return nil, err
	}

	return &TournamentStorage{Db: db}, nil
}
func (tournamentdb *TournamentStorage) CreateTournament(creatorID, name, description string) (string, error) {
	return "nil", nil
}

// GET ALL TOURNAMENT
func (tournamentdb *TournamentStorage) GetAllTournaments() ([]models.Tournament, error) {
	rows, err := tournamentdb.Db.Query(`SELECT tournament_id, creator_id, name, description, created_at FROM tournaments`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tournaments []models.Tournament

	for rows.Next() {
		var t models.Tournament
		err := rows.Scan(&t.ID, &t.CreatorID, &t.Name, &t.Description, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		tournaments = append(tournaments, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tournaments, nil
}

func (tournamentdb *TournamentStorage) AddProblem(tournamentID string, problem models.Problem) (bool, error) {
	return false, nil
}
func (tournamentdb *TournamentStorage) GetProblems(tournamentID string) ([]models.GetProblemResponse, error) {
	row, err := tournamentdb.Db.Query(`SELECT id, title , description, maxScore WHERE tournament_id == ?`, tournamentID)
	if err != nil {

	}
	var problems []models.GetProblemResponse
	for row.Next() {
		var prob models.GetProblemResponse
		err := row.Scan(&prob.Id, &prob.Title, &prob.Description, &prob.MaxScore)
		if err != nil {

		}
		problems = append(problems, prob)
	}
	return problems, nil
}

func (tournamentdb *TournamentStorage) AddTestCase(problemID string, testCase models.TestCase) (bool, error) {
	//jwt token 
}
func (tournamentdb *TournamentStorage) AddParticipant(tournamentID) (bool, error) {
	//jwt token
}
func (tournamentdb *TournamentStorage) GetAllParticipants(tournamentID string) ([]models.GetTournamentParticipantResponse, error) {
	row, err := tournamentdb.Db.Query(`SELECT user_id, user_name, score from tournament_participants WHERE tournament_id=?`, tourtournamentID)

	if err != nil {

	}
	var tournametParticipant []models.GetTournamentParticipantResponse
	for row.Next() {
		var t models.GetTournamentParticipantResponse
		err = row.Scan(&t.UserID, &t.UserName, &t.Score)
		if err != nil {

		}
		tournametParticipant = append(tournametParticipant, t)
	}
	return tournametParticipant,nil
}

tournament_id TEXT NOT NULL,
		user_id TEXT NOT NULL,
		user_name TEXT NOT NULL,
		joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		score INTEGER DEFAULT 0,