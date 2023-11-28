package sneakers

import (
	"backend/internal/handlers"
	"backend/internal/users"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

const (
	sneakersURL = users.API + "/sneakers"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(sneakersURL, h.Test)
}

func (h *handler) Test(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	log.Println("all work")
}
