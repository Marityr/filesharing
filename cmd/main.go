package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Marityr/filesharing"
	"github.com/Marityr/filesharing/pkg/handler"
	"github.com/joho/godotenv"
)

func main() {
	initConfig()
	srv := new(filesharing.Server)
	runHttp(srv)
}

func runHttp(srv *filesharing.Server) {
	go func() {
		if err := srv.Run(os.Getenv("PORT"), handler.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Println("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("App Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
