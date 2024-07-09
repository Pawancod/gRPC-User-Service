package server

import (
	"log"
	"net"

	"github.com/Pawancod/gRPC-User-Service/internal/repository"
	"github.com/Pawancod/gRPC-User-Service/internal/service"
	"github.com/Pawancod/gRPC-User-Service/proto/github.com/Pawancod/gRPC-User-Service/proto"
	"google.golang.org/grpc"
)

func StartGRPCServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	repo := repository.NewUserRepository()
	userService := service.NewUserService(repo)
	proto.RegisterUserServiceServer(grpcServer, userService)

	log.Printf("Server is running on port %s...", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
