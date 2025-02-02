package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"payments/app"
	"payments/config"
)

const configFile = "config.yml"

func main() {
	cfg, err := config.Load(configFile)
	if err != nil {
		log.Fatal("unable to load config file")
	}

	newApp, _ := app.NewApp(cfg)
	err = newApp.StartServer()
	if err != nil {
		log.Fatal("Error starting app", err)
	}

	signl := make(chan os.Signal, 1)

	signal.Notify(signl, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-signl

	err = newApp.ShutdownServer()
	if err != nil {
		log.Fatal("Error shutting down serving", err)
	}
}
