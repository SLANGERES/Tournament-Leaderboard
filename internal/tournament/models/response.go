package models

type GetProblemResponse struct {
	Id          string
	Title       string
	Description string
	MaxScore    int
}
type GetTournamentParticipantResponse struct {
	UserID   string `json:"user_id"`
	UserName string `json:"username"` // SK
	Score    int    `json:"score"`    // Starts at 0
}
