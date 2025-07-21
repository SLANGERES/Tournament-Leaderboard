package models

import "time"

// CreateTournament represents the structure for creating a tournament
type Tournament struct {
	ID          string    `json:"tournament_id"`
	CreatorID   string    `json:"creator_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// Problem represents a single coding problem/question
type Problem struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	InputFormat  string   `json:"input_format"`
	OutputFormat string   `json:"output_format"`
	TestCase     TestCase `json:"test_case"`
	MaxScore     int      `json:"max_score"`
}

// TestCase represents a sample test case for a problem
type TestCase struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

// TournamentParticipant represents a user participating in a tournament
type TournamentParticipant struct {
	TournamentID string `json:"tournament_id"` // PK
	UserID       string `json:"user_id"`       // SK
	Score        int    `json:"score"`         // Starts at 0
}


