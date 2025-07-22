package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/SLANGERES/Tournament-Lederboard/internal/common/jwt"
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

func LogInUser(db *repository.UserStorage, jwt *jwt.JwtMaker) http.HandlerFunc {
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

		id, err := db.LoginUser(userLoginData.UserName, userLoginData.Password)
		if err != nil {
			util.HttpError(w, http.StatusUnauthorized, fmt.Errorf("invalid username or password"))
			return
		}
		strId := strconv.FormatInt(id, 10)

		//JWT
		token, err := jwt.GenerateToken(strId, userLoginData.UserName, "user")
		if err != nil {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("unable to generate"))
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:   "user-access-token",
			Value:  token,
			MaxAge: int((30 * time.Minute).Seconds()),
		})

		util.HttpResponse(w, http.StatusOK, "Login successful")
	}
}
