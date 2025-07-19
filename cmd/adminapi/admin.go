package main

//Admin Backend API

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/SLANGERES/Tournament-Lederboard/config"
	"github.com/SLANGERES/Tournament-Lederboard/internal/admin/handler"
	"github.com/SLANGERES/Tournament-Lederboard/internal/admin/repository"
)

func main() {

	cnf, err := config.SetConfig()
	if err != nil {
		slog.Warn("unable to get config file")
		os.Exit(1)
	}
	router := http.NewServeMux()

	db, err := repository.ConfigAdminDB(cnf.AdminDB)
	if err != nil {
		slog.Info("Unable to connect to the Admin DB" + err.Error())
	}
	//! Routers Endpoints
	router.HandleFunc("POST /sign-up", handler.Signup(db))
	router.HandleFunc("POST /log-in", handler.Login(db))

	slog.Info("Admin Server started at" + cnf.HttpServer.AdminAddress)
	err = http.ListenAndServe(
		cnf.HttpServer.AdminAddress,
		router,
	)
	if err != nil {
		slog.Info("Server start fail !")
	}

}
