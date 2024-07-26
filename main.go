package main

import (
	"fmt"
	"kindred/api"
	"log"
)

const (
	serverAddress = "0.0.0.0"
)

func main() {
	server := api.NewServer()
	err := server.Start(serverAddress + ":" + "8080")
	if err != nil {
		log.Fatal("Server start error", err)
	}
	fmt.Println("server started")
}
