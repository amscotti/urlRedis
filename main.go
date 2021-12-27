package main // github.com:amscotti/urlRedis

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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

func gracefulShutdown(app *fiber.App) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	log.Println("Gracefully shutting down...")
	if err := app.Shutdown(); err != nil {
		return err
	}

	log.Println("Running cleanup tasks...")
	if err := storage.DBConn.Close(); err != nil {
		return err
	}

	log.Println("Fiber was successful shutdown.")
	return nil
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

	errs := make(chan error)
	go func() {
		if err := app.Listen(":" + port); err != nil {
			errs <- err
		}
	}()
	go func() { errs <- gracefulShutdown(app) }()

	err := <-errs
	if err != nil {
		log.Fatal(err)
	}
}
