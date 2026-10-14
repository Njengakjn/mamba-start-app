package handlers

import (
	"net/http"
	"time"

	"github.com/kjnjenga/mamba"
	"github.com/kjnjenga/webapp/models"
)

type Handlers struct {
	mamba  *mamba.Mamba
	models models.Models
}

func New(m *mamba.Mamba, models models.Models) *Handlers {
	return &Handlers{
		mamba:  m,
		models: models,
	}
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	defer h.mamba.LoadTime(time.Now())
	err := h.mamba.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.mamba.ErrorLog.Println("error rendering page: ", err)
	}
}
