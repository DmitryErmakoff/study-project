package sneakers

import (
	"backend/internal/handlers"
	"backend/internal/users"
	"backend/pkg/client/postgresql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	sneakersURL        = users.API + "/sneakers"
	favoritesURL       = users.API + "/favorites"
	addFavoritesURL    = users.API + "/addfavorites"
	deleteFavoritesURL = users.API + "/deletefavorites"
	cartURL            = users.API + "/cart"
	addCartURL         = users.API + "/addcart"
	deleteCartURL      = users.API + "/deletecart"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(sneakersURL, h.GetAllSnekaers)
	router.POST(favoritesURL, h.GetAllFavorites)
	router.POST(addFavoritesURL, h.AddToFavorites)
	router.POST(deleteFavoritesURL, h.DeleteFromFavorites)
	router.POST(cartURL, h.GetAllCart)
	router.POST(addCartURL, h.AddToCart)
	router.POST(deleteCartURL, h.DeleteFromCart)
}

func (h *handler) GetAllSnekaers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var sneakers []Sneaker
	rows, err := postgresql.DB.Query(`SELECT 
    s.id_sneaker,
    s.name_sneaker,
    s.brand_sneaker,
    s.description_sneaker,
    sc.color,
    i.first_photo,
    i.second_photo,
    i.third_photo,
    i.fourth_photo,
	s.price,
	s.date_release
FROM
    sneakers AS s
    INNER JOIN sneakers_color AS sc ON s.id_sneaker = sc.id_sneaker
    INNER JOIN images AS i ON s.id_images = i.id_images;`)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	for rows.Next() {
		var sneaker Sneaker
		var first_photo string
		var second_photo string
		var third_photo string
		var thourt_photo string
		err = rows.Scan(&sneaker.Id, &sneaker.Model, &sneaker.Brand, &sneaker.Description, &sneaker.Color, &first_photo, &second_photo, &third_photo, &thourt_photo, &sneaker.Price, &sneaker.Date_release)
		sneaker.Images = append(sneaker.Images, first_photo, second_photo, third_photo, thourt_photo)
		sneakers = append(sneakers, sneaker)
	}

	jsonResponse, err := json.Marshal(sneakers)
	w.WriteHeader(200)
	w.Write(jsonResponse)
}

func (h *handler) GetAllFavorites(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var sneakers []Sneaker
	var request FavoriteRequest
	json.NewDecoder(r.Body).Decode(&request)
	rows, err := postgresql.DB.Query(`SELECT
    s.id_sneaker,
    s.name_sneaker,
    s.brand_sneaker,
    s.description_sneaker,
    sc.color,
    i.first_photo,
    i.second_photo,
    i.third_photo,
    i.fourth_photo,
    s.price,
    s.date_release
FROM
    favorites_sneaker AS f
    INNER JOIN sneakers AS s ON f.id_sneaker = s.id_sneaker
    INNER JOIN sneakers_color AS sc ON s.id_sneaker = sc.id_sneaker
    INNER JOIN images AS i ON s.id_images = i.id_images
WHERE
    f.id_user = $1;`, request.Id_user)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	for rows.Next() {
		var sneaker Sneaker
		var first_photo string
		var second_photo string
		var third_photo string
		var thourt_photo string
		err = rows.Scan(&sneaker.Id, &sneaker.Model, &sneaker.Brand, &sneaker.Description, &sneaker.Color, &first_photo, &second_photo, &third_photo, &thourt_photo, &sneaker.Price, &sneaker.Date_release)
		sneaker.Images = append(sneaker.Images, first_photo, second_photo, third_photo, thourt_photo)
		sneakers = append(sneakers, sneaker)
	}
	jsonResponse, err := json.Marshal(sneakers)
	w.WriteHeader(200)
	w.Write(jsonResponse)
}

func (h *handler) AddToFavorites(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var request FavoriteRequest
	json.NewDecoder(r.Body).Decode(&request)
	insertFavoriteSneaker := `INSERT INTO favorites_sneaker(id_user, id_sneaker)
		VALUES($1, $2)`
	_, err := postgresql.DB.Exec(insertFavoriteSneaker, users.UserService.Id, request.Id_sneaker)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(200)
}

func (h *handler) DeleteFromFavorites(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var request FavoriteRequest
	json.NewDecoder(r.Body).Decode(&request)
	deleteFavoriteSneaker := `DELETE FROM favorites_sneaker WHERE id_user = $1 AND id_sneaker = $2`
	_, err := postgresql.DB.Exec(deleteFavoriteSneaker, users.UserService.Id, request.Id_sneaker)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(200)
}

func (h *handler) GetAllCart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var sneakers []Sneaker
	var request FavoriteRequest
	json.NewDecoder(r.Body).Decode(&request)
	rows, err := postgresql.DB.Query(`SELECT
    s.id_sneaker,
    s.name_sneaker,
    s.brand_sneaker,
    s.description_sneaker,
    sc.color,
    i.first_photo,
    i.second_photo,
    i.third_photo,
    i.fourth_photo,
    s.price,
    s.date_release
FROM
    user_cart AS u
    INNER JOIN sneakers AS s ON u.id_sneaker = s.id_sneaker
    INNER JOIN sneakers_color AS sc ON s.id_sneaker = sc.id_sneaker
    INNER JOIN images AS i ON s.id_images = i.id_images
WHERE
    u.id_user = $1;`, request.Id_user)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	for rows.Next() {
		var sneaker Sneaker
		var first_photo string
		var second_photo string
		var third_photo string
		var thourt_photo string
		err = rows.Scan(&sneaker.Id, &sneaker.Model, &sneaker.Brand, &sneaker.Description, &sneaker.Color, &first_photo, &second_photo, &third_photo, &thourt_photo, &sneaker.Price, &sneaker.Date_release)
		sneaker.Images = append(sneaker.Images, first_photo, second_photo, third_photo, thourt_photo)
		sneakers = append(sneakers, sneaker)
	}
	jsonResponse, err := json.Marshal(sneakers)
	w.WriteHeader(200)
	w.Write(jsonResponse)
}

func (h *handler) AddToCart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var request FavoriteRequest
	json.NewDecoder(r.Body).Decode(&request)
	insertFavoriteSneaker := `INSERT INTO user_cart(id_user, id_sneaker)
		VALUES($1, $2)`
	_, err := postgresql.DB.Exec(insertFavoriteSneaker, users.UserService.Id, request.Id_sneaker)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(200)
}

func (h *handler) DeleteFromCart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var request FavoriteRequest
	json.NewDecoder(r.Body).Decode(&request)
	deleteFavoriteSneaker := `DELETE FROM user_cart WHERE id_user = $1 AND id_sneaker = $2`
	_, err := postgresql.DB.Exec(deleteFavoriteSneaker, users.UserService.Id, request.Id_sneaker)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(200)
}
