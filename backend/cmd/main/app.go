package main

import (
	"backend/internal/sneakers"
	"backend/internal/users"
	"github.com/gookit/slog"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	slog.Info("CREATE ROUTER")
	router := httprouter.New()

	slog.Info("REGISTER USER HANDLER")
	handler := users.NewHandler()
	handler.Register(router)
	slog.Info("REGISTER SNEAKER HANDLER")
	handler = sneakers.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	slog.Info("START APPLICATION")

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	slog.Info("SERVER IS LISTENING PORT 0.0.0.0:1234")
	log.Fatalln(server.Serve(listener))
}
