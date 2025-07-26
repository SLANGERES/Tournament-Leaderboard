package main

//Admin Backend API

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/SLANGERES/Tournament-Lederboard/config"
	"github.com/SLANGERES/Tournament-Lederboard/internal/admin/handler"
	"github.com/SLANGERES/Tournament-Lederboard/internal/admin/repository"
	"github.com/SLANGERES/Tournament-Lederboard/internal/admin/service"
	"github.com/SLANGERES/Tournament-Lederboard/internal/common/jwt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	serviceDB := service.NewUserService(db)
	jwtMaker := jwt.NewJwtMaker(cnf.JwtKey)

	//!Prometheus Server Metrices
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":9091", nil) // Different port for Prometheus scraping
	}()

	//! Routers Endpoints
	router.HandleFunc("POST /v1/admins/signup", handler.Signup(*serviceDB))
	router.HandleFunc("POST /v1/admins/login", handler.Login(*serviceDB, jwtMaker))

	slog.Info("Admin Server started at" + cnf.HttpServer.AdminAddress)
	err = http.ListenAndServe(
		cnf.HttpServer.AdminAddress,
		router,
	)
	if err != nil {
		slog.Info("Server start fail !")
	}

}
