package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/amscotti/urlRedis/handlers"
	"github.com/amscotti/urlRedis/storage"
	"github.com/go-zoo/bone"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	db := storage.NewRedis()

	router := bone.New()
	router.Post("/", http.HandlerFunc(handlers.CreateKey(db)))
	router.Get("/:key", http.HandlerFunc(handlers.RedirectKey(db)))
	router.Get("/get/:key", http.HandlerFunc(handlers.GetKey(db)))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Service started on port %s", port)
	http.ListenAndServe(":"+port, router)
}
