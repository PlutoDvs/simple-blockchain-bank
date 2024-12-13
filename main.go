package main

import (
	"github.com/PlutoDvs/Gochain/loan"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Args[1]
	router := loan.NewRouter(port)

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST"})

	log.Fatal(http.ListenAndServe(":"+port,
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
