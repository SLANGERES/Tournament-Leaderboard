package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/SLANGERES/Tournament-Lederboard/internal/common/util"
	"github.com/SLANGERES/Tournament-Lederboard/internal/user/models"
	"github.com/SLANGERES/Tournament-Lederboard/internal/user/repository"
	"github.com/go-playground/validator/v10"
)

func SignInUser(db *repository.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userData models.CreateUser
		if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
			util.HttpError(w, http.StatusBadRequest, fmt.Errorf("unable to decode"))
			slog.Info("error in decode")
			return
		}

		if err := validator.New().Struct(userData); err != nil {
			util.HttpError(w, http.StatusBadRequest, fmt.Errorf("validation error"))
			slog.Info("error in validation")
			return
		}

		_, err := db.CreateUser(userData.UserName, userData.Email, userData.Password)
		if err != nil {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("failed to create user"))
			return
		}

		util.HttpResponse(w, http.StatusCreated, "User created successfully")
	}
}

func LogInUser(db *repository.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userLoginData models.LoginUser

		if err := json.NewDecoder(r.Body).Decode(&userLoginData); err != nil {
			util.HttpError(w, http.StatusBadRequest, err)
			return
		}

		if err := validator.New().Struct(userLoginData); err != nil {
			util.HttpError(w, http.StatusBadRequest, fmt.Errorf("validation error"))
			return
		}

		success, err := db.LoginUser(userLoginData.UserName, userLoginData.Password)
		if err != nil || !success {
			util.HttpError(w, http.StatusUnauthorized, fmt.Errorf("invalid username or password"))
			return
		}

		// TODO: Replace with actual JWT generation
		util.HttpResponse(w, http.StatusOK, "Login successful")
	}
}
