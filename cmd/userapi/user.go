package main

import (
	"log/slog"
	"net/http"

	"github.com/SLANGERES/Tournament-Lederboard/config"
	"github.com/SLANGERES/Tournament-Lederboard/internal/common/jwt"
	"github.com/SLANGERES/Tournament-Lederboard/internal/user/handler"
	"github.com/SLANGERES/Tournament-Lederboard/internal/user/repository"
	"github.com/SLANGERES/Tournament-Lederboard/internal/user/service"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	cnf, err := config.SetConfig()
	if err != nil {
		slog.Warn("Unable to get the config file")
	}
	jwtMaker := jwt.NewJwtMaker(cnf.JwtKey)
	router := http.NewServeMux()

	db, err := repository.ConfigUserStorage(cnf.UserDB)
	service := service.NewUserService(db)

	if err != nil {
		slog.Error(err.Error())
	}

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":9092", nil) // Different port for Prometheus scraping
	}()
	//! Routers Endpoints
	router.HandleFunc("POST /v1/user/signup", handler.SignInUser(*service))
	router.HandleFunc("POST /v1/user/login", handler.LogInUser(*service, jwtMaker))

	slog.Info("Server started at " + cnf.HttpServer.UserAddress)
	err = http.ListenAndServe(
		cnf.HttpServer.UserAddress,
		router,
	)
	if err != nil {
		slog.Info("User Server start fail !")
	}

}
