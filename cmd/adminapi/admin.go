package main

//Admin Backend API

import (
	"log/slog"
	"net/http"

	"github.com/SLANGERES/Tournament-Lederboard/config"
	"github.com/SLANGERES/Tournament-Lederboard/internal/admin/handler"
	"github.com/SLANGERES/Tournament-Lederboard/internal/admin/repository"
)

func main() {

	cnf := config.SetConfig()
	router := http.NewServeMux()

	db, err := repository.ConfigAdminDB(cnf.AdminDB)
	if err != nil {
		slog.Info("Unable to connect to the Admin DB" + err.Error())
	}
	//! Routers Endpoints
	router.HandleFunc("POST /sign-up", handler.Signup(db))
	router.HandleFunc("POST /log-in", handler.Login(db))

	err = http.ListenAndServe(
		cnf.HttpServer.AdminAddress,
		router,
	)
	if err != nil {
		slog.Info("Server start fail !")
	}
	slog.Info("Admin Server started at" + cnf.HttpServer.AdminAddress)

}
