package models

type CreateUser struct {
	Id       string `json:"id"`
	Email    string `json:"email" validate:"required"`
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginUser struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
