package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"restApiSneakerShop/internal/handlers/sneaker"
	"time"
)

func main() {
	log.Println("create router")
	router := httprouter.New()

	log.Println("register user handler")
	handler := sneaker.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	log.Println("start application")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Println("server is listening port 1234")
	log.Fatalln(server.Serve(listener))
}
