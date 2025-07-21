package main

import (
	"log"

	"github.com/ggualbertosouza/game/api"
	"github.com/ggualbertosouza/game/internal/config"
)

func main() {
	startServer()
}

func startServer() {
	config.LoadEnv()

	port := config.GetEnv("SERVER_PORT", ":8080")
	server := api.New(port)

	log.Fatal(server.Start())
}