package main

import (
	"log"
	"os"

	"github.com/kjnjenga/mamba"
	"github.com/kjnjenga/startapp/handlers"
	"github.com/kjnjenga/startapp/middleware"
	"github.com/kjnjenga/startapp/models"
)

type application struct {
	mamba      *mamba.Mamba
	handlers   *handlers.Handlers
	models     models.Models
	middleware *middleware.Middleware
}

func main() {
	app := initialize()

	app.mamba.ListenAndServe()
}

func initialize() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	m := &mamba.Mamba{}

	err = m.New(path)
	if err != nil {
		log.Fatal(err)
	}

	m.Debug = true
	m.Title = "Webapp"

	models := models.New(m.DB.Pool)
	handlers := handlers.New(m, models)
	middleware := middleware.New(m, models)

	app := &application{
		mamba:      m,
		handlers:   handlers,
		models:     models,
		middleware: middleware,
	}

	app.mamba.Routes = app.routes()

	return app
}
