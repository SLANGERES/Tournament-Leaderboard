package main

import (
	"log/slog"
	"net/http"

	"github.com/SLANGERES/Tournament-Lederboard/config"
	"github.com/SLANGERES/Tournament-Lederboard/internal/user/handler"
	"github.com/SLANGERES/Tournament-Lederboard/internal/user/repository"
)

func main() {

	cnf, err := config.SetConfig()
	if err != nil {
		slog.Warn("Unable to get the config file")
	}

	router := http.NewServeMux()

	db, err := repository.ConfigUserStorage(cnf.UserDB)

	if err != nil {
		slog.Error(err.Error())
	}

	//! Routers Endpoints
	router.HandleFunc("POST /v1/user/sign-up", handler.SignInUser(db))
	router.HandleFunc("POST /v1/user/log-in", handler.LogInUser(db))

	slog.Info("Server started at " + cnf.HttpServer.UserAddress)
	err = http.ListenAndServe(
		cnf.HttpServer.UserAddress,
		router,
	)
	if err != nil {
		slog.Info("User Server start fail !")
	}

}
