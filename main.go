package main // github.com:amscotti/urlRedis

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/amscotti/urlRedis/handlers"
	"github.com/amscotti/urlRedis/storage"
)

func initDatabase() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = ":6379"
	}
	storage.DBConn = storage.NewRedis(redisURL)
}

func setUpRoutes(app *fiber.App) {
	keyRoutes := app.Group("/v1/keys")
	keyRoutes.Post("/", handlers.CreateKey)
	keyRoutes.Get(":key", handlers.GetKey)

	app.Get("/:key", handlers.RedirectKey)
}

func gracefulShutdown(app *fiber.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	initDatabase()
	setUpRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	go func() { gracefulShutdown(app) }()

	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
