package users

import (
	"backend/internal/handlers"
	"backend/pkg/client/postgresql"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

const (
	API         = "/api"
	registerURL = API + "/register"
	loginURL    = API + "/login"
)

var UserService User

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.POST(loginURL, h.LoginUser)
	router.POST(registerURL, h.RegisterUser)
}

func (h *handler) RegisterUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	insertUser := `insert into "users"("name", "password", "access_rights") values($1, $2, $3)`
	_, err := postgresql.DB.Exec(insertUser, u.Name, u.Password, "client")
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Пользователь %s, с паролем %s, правами доступа %s \nУспешно зарегистрирован", u.Name, u.Password, "client")))

}

func (h *handler) LoginUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	rows, err := postgresql.DB.Query(`SELECT * FROM users`)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	for rows.Next() {
		var id int
		var name string
		var password string
		var permission string

		err = rows.Scan(&id, &name, &password, &permission)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		if u.Name == name && u.Password == password {
			UserService.Id = id
			UserService.Name = name
			UserService.Password = password
			UserService.Permission = permission
			jsonResponse, _ := json.Marshal(UserService)
			w.WriteHeader(200)
			w.Write(jsonResponse)
			log.Println(fmt.Sprintf("Вход выполнен:\nПользователь: %s\nПрава доступа: %s", UserService.Name, UserService.Permission))
			return

		}
	}
	w.WriteHeader(400)
	w.Write([]byte("Неправильное имя пользователя или пароль"))
	return
}
