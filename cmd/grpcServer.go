package cmd

import (
	"fmt"
	"github.com/sinojin/questioner/internal/common/genproto/questioner"
	"github.com/sinojin/questioner/internal/questioner/adapters"
	"github.com/sinojin/questioner/internal/questioner/ports"
	"github.com/sinojin/questioner/internal/questioner/services"
	"google.golang.org/grpc"
	"net"
)

func GrpcServer() {
	repoq := adapters.NewQuestionRepository()
	repos := adapters.NewStatisticsRepository()

	questionerService := services.NewQuestionService(repoq, repos)

	grpcServerImplementation := ports.NewGrpcServer(questionerService)

	//basic serve code for grpc
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	questioner.RegisterQuestionerServiceServer(grpcServer, grpcServerImplementation)
	fmt.Println("Server Started.")
	go grpcServer.Serve(lis)
}
