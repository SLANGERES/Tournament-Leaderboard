package models

import "time"

type CreateTournamet struct {
	Id          string     `json:"tournament_id"`
	CreaterID   string     `json:"creater_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Question    []Question `json:"question"`
}

type Question struct {
	Questions string
	Point     string
}

type TournamentParticipant struct {
	TournamentID string    `json:"tournament_id"` // PK
	UserID       string    `json:"user_id"`       // SK
	JoinedAt     time.Time `json:"joined_at"`
	Score        int       `json:"score"` // Initialize with 0
}
