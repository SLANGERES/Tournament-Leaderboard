package handler

import (
	"net/http"

	"github.com/SLANGERES/Tournament-Lederboard/internal/admin/repository"
)

func Signup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}

func Login(Storage *repository.DbConnection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
