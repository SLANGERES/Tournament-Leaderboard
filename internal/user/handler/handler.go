package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SLANGERES/Tournament-Lederboard/internal/common/util"
	"github.com/SLANGERES/Tournament-Lederboard/internal/user/models"
	"github.com/SLANGERES/Tournament-Lederboard/internal/user/repository"
	"github.com/go-playground/validator/v10"

)

func SignInUser(db repository.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userData models.CreateUser
		err := json.NewDecoder(r.Body).Decode(&userData)
		if err != nil {
			util.HttpError(w, http.StatusInternalServerError, err)
		}
		err = validator.New().Struct(userData)
		if err != nil {
			util.HttpError(w, http.StatusInternalServerError, err)
		}

		_, err = db.CreateAdmin(userData.UserName, userData.Email, userData.Password)
		if err != nil {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("database error"))
		}
		util.HttpResponse(w, http.StatusCreated, "User created Successfully")
	}

}
func LogInUser(db repository.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userLoginData models.LoginUser

		err := json.NewDecoder(r.Body).Decode(&userLoginData)

		if err != nil {
			util.HttpError(w, http.StatusInternalServerError, err)
		}
		err = validator.New().Struct(userLoginData)

		if err != nil {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("validation error"))
		}
		_, err = db.LoginUser(userLoginData.UserName, userLoginData.Password)

		if err != nil {
			util.HttpError(w, http.StatusInternalServerError, err)
		}
		//Todo Jwt Auth ->

		util.HttpResponse(w, http.StatusInternalServerError, "Login successfully")

	}
}
