package middleware

import (
	"github.com/kjnjenga/mamba"
	"github.com/kjnjenga/startapp/models"
)

type Middleware struct {
	mamba  *mamba.Mamba
	models models.Models
}

func New(mamba *mamba.Mamba, models models.Models) *Middleware {
	return &Middleware{
		mamba:  mamba,
		models: models,
	}
}
