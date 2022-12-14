package main

import (
	"github.com/DadaxonOlimov/api"
	"github.com/DadaxonOlimov/services"
)

func main() {

	grpcClients, _ := services.NewGrpcClients()

	server := api.NewServer(grpcClients)

	err := server.Run(":8081")

	if err != nil {
		return
	}
}