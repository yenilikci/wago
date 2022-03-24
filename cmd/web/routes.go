package main

import (
	"github.com/bmizerany/pat"
	"github.com/yenilikci/wago/pkg/config"
	"github.com/yenilikci/wago/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
