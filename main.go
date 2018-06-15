package main // github.com:amscotti/urlRedis

import (
	"log"
	"net/http"
	"os"

	"github.com/amscotti/urlRedis/handlers"
	"github.com/amscotti/urlRedis/storage"

	"github.com/gorilla/mux"
)

func main() {
	db := storage.NewRedis()

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.CreateKey(db)).Methods("POST")
	router.HandleFunc("/{key}", handlers.RedirectKey(db)).Methods("GET")
	router.HandleFunc("/get/{key}", handlers.GetKey(db)).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("Service started on port %s", port)
	http.ListenAndServe(":"+port, router)
}
