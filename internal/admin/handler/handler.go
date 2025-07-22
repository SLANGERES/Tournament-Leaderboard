package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/SLANGERES/Tournament-Lederboard/internal/admin/models"
	"github.com/SLANGERES/Tournament-Lederboard/internal/admin/repository"
	"github.com/SLANGERES/Tournament-Lederboard/internal/common/jwt"
	"github.com/SLANGERES/Tournament-Lederboard/internal/common/util"
	"github.com/go-playground/validator/v10"
)

func Signup(storage *repository.DbConnection) http.HandlerFunc {
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

		if _, err := storage.CreateAdmin(newAdmin.Email, newAdmin.UserName, newAdmin.Password); err != nil {
			util.HttpError(w, http.StatusInternalServerError, fmt.Errorf("failed to create admin: %v", err))
			return
		}

		util.HttpResponse(w, http.StatusCreated, "Admin created successfully")
	}
}

func Login(Storage *repository.DbConnection, jwt *jwt.JwtMaker) http.HandlerFunc {
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
