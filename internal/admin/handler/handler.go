package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/SLANGERES/Tournament-Lederboard/internal/admin/models"
	"github.com/SLANGERES/Tournament-Lederboard/internal/admin/service"
	"github.com/SLANGERES/Tournament-Lederboard/internal/common/jwt"
	"github.com/SLANGERES/Tournament-Lederboard/internal/common/util"
	"github.com/go-playground/validator/v10"
)

// Signup godoc
// @Summary Admin signup
// @Description Register a new admin account
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin body models.CreateAdmin true "Admin Signup Data"
// @Success 201 {string} string "Admin created successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /admin/signup [post]
func Signup(storage service.AdminService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newAdmin models.CreateAdmin

		if err := json.NewDecoder(r.Body).Decode(&newAdmin); err != nil {
			util.HttpError(w, http.StatusBadRequest, fmt.Errorf("unable to parse request body"))
			return
		}

		validate := validator.New()
		if err := validate.Struct(newAdmin); err != nil {
			util.HttpError(w, http.StatusBadRequest, fmt.Errorf("validation error: %v", err))
			return
		}

		if _, err := storage.SignupAdmin(newAdmin.Email, newAdmin.UserName, newAdmin.Password); err != nil {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("failed to create admin: %v", err))
			return
		}

		util.HttpResponse(w, http.StatusCreated, "Admin created successfully")
	}
}

// Login godoc
// @Summary Admin login
// @Description Authenticate an admin and return JWT in cookie
// @Tags Admin
// @Accept json
// @Produce json
// @Param credentials body models.LogInAdmin true "Admin Login Data"
// @Success 200 {string} string "Login successful"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /admin/login [post]
func Login(Storage service.AdminService, jwt *jwt.JwtMaker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var adminLogin models.LogInAdmin
		err := json.NewDecoder(r.Body).Decode(&adminLogin)

		if err != nil {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("unable to decode request "+err.Error()))
			return
		}

		err = validator.New().Struct(adminLogin)

		if err != nil {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("validation error"+err.Error()))
			return
		}

		id, err := Storage.LoginAdmin(adminLogin.UserName, adminLogin.Password)
		if err != nil {
			util.HttpError(w, http.StatusInternalServerError, err)
			return
		}
		strid := strconv.FormatInt(id, 10)

		token, err := jwt.GenerateToken(strid, adminLogin.UserName, "admin")
		if err != nil {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("unable to generate jwt token"))
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:   "admin-access-token",
			Value:  token,
			MaxAge: int((30 * time.Minute).Seconds()),
		})

		util.HttpResponse(w, http.StatusOK, "Login sucessfull")
	}
}
