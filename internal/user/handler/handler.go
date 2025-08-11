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
	"github.com/SLANGERES/Tournament-Lederboard/internal/user/service"
	"github.com/go-playground/validator/v10"
)

// Signup godoc
// @Summary User signup
// @Description Register a new User account
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.CreateUser true "User Signup Data"
// @Success 201 {string} string "User created successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /user/signup [post]
func SignInUser(servie service.UserService) http.HandlerFunc {
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

		_, err := servie.Signup(userData.UserName, userData.Email, userData.Password)
		if err != nil {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("failed to create user"))
			return
		}

		util.HttpResponse(w, http.StatusCreated, "User created successfully")
	}
}

// Login godoc
// @Summary User login
// @Description Authenticate an user and return JWT in cookie
// @Tags User
// @Accept json
// @Produce json
// @Param credentials body models.LoginUser true "User Login Data"
// @Success 200 {string} string "Login successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /user/login [post]
func LogInUser(servie service.UserService, jwt *jwt.JwtMaker) http.HandlerFunc {
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
		id, err := servie.Login(userLoginData.UserName, userLoginData.Password)
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
