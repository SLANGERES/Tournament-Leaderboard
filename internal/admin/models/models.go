package models

type CreateAdmin struct {
	Id       string `json:"id"`
	Email    string `json:"email" validate:"required"`
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LogInAdmin struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
