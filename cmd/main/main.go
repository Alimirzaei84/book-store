package main

import (
	"log"
	"net/http"

	"github.com/Alimirzaei84/book-store/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
    routes.RegisterBookStoreRoutes(router)
    http.Handle("/", router)
    log.Fatal(http.ListenAndServe(":8080", nil))
}