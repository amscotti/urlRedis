package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/amscotti/urlRedis/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/amscotti/urlRedis/handlers"
	"github.com/amscotti/urlRedis/storage"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	db := storage.NewRedis()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handlers.CreateKey(db)).Methods("POST")
	router.HandleFunc("/{key}", handlers.RedirectKey(db)).Methods("GET")
	router.HandleFunc("/get/{key}", handlers.GetKey(db)).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}
