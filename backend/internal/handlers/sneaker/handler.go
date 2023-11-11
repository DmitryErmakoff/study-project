package sneaker

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"restApiSneakerShop/internal/handlers"
)

const (
	sneakersURL = "/sneakers"
	sneakerURL  = "/sneakers/:id"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(sneakersURL, h.GetListSneakers)
	router.GET(sneakerURL, h.GetSneakerByID)
	router.POST(sneakerURL, h.CreateSneaker)
	router.PUT(sneakerURL, h.UpdateSneaker)
	router.PATCH(sneakerURL, h.PartiallyUpdateSneaker)
	router.DELETE(sneakerURL, h.DeleteSneakerByID)
}

func (h *handler) GetListSneakers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("this is list of sneakers"))
}

func (h *handler) GetSneakerByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("this is getsneaker by id"))
}

func (h *handler) CreateSneaker(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("this is create sneaker"))
}

func (h *handler) UpdateSneaker(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is update by id"))
}

func (h *handler) PartiallyUpdateSneaker(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is partially update by id"))
}

func (h *handler) DeleteSneakerByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is delete by id"))
}
