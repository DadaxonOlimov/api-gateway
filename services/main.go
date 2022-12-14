package services

import (
	"github.com/DadaxonOlimov/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	TodosService() proto.TodoServiceClient
}

type grpcClients struct {
	todosService proto.TodoServiceClient
}

func NewGrpcClients() (ServiceManager, error) {
	todosService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", "localhost", 8082),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		todosService: proto.NewTodoServiceClient(todosService),
	}, nil
}

func (g grpcClients) TodosService() proto.TodoServiceClient {
	return g.todosService
}