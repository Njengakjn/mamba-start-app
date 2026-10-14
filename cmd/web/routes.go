package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// middleware must come before any routes
	a.mamba.Routes.Use(a.middleware.CheckRemember)

	// add routes here
	a.mamba.Routes.Get("/", a.handlers.Home)

	// static routes
	fileserver := http.FileServer(http.Dir("./public"))
	a.mamba.Routes.Handle("/public/*", http.StripPrefix("/public", fileserver))

	return a.mamba.Routes
}
