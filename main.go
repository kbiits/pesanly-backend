package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"github.com/kbiits/chat-app/handler"
	"github.com/kbiits/chat-app/repo"
	"github.com/kbiits/chat-app/usecase"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	dsn := os.Getenv("dsn")

	err := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatalf("failed init mgm, err : %v\n", err)
	}
}

func main() {

	repo := repo.NewRepo()
	uc := usecase.NewUsecase(repo)
	h := handler.NewHandler(uc)

	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	initializeRoutes(app, h)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	err := app.Listen(":8080")
	if err != nil {
		log.Fatalf("err listen : %v\n", err)
	}
	fmt.Println("Running cleanup tasks...")
}
