package main

import (
	"context"
	"log"

	"github.com/Pawancod/gRPC-User-Service/proto/github.com/Pawancod/gRPC-User-Service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up credentials, insecure for local development

	// Set up a connection to the server.
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client instance
	client := proto.NewUserServiceClient(conn)

	// Example: Calling GetUser
	getUserResponse, err := client.GetUser(context.Background(), &proto.GetUserRequest{Id: 1})
	if err != nil {
		log.Fatalf("failed to get user: %v", err)
	}
	log.Printf("User: %v", getUserResponse.GetUser())

	// Example: Calling GetUsers
	getUsersResponse, err := client.GetUsers(context.Background(), &proto.GetUsersRequest{Ids: []int32{1, 2}})
	if err != nil {
		log.Fatalf("failed to get users: %v", err)
	}
	log.Printf("Users: %v", getUsersResponse.GetUsers())

	// Example: Calling SearchUsers
	searchUsersResponse, err := client.SearchUsers(context.Background(), &proto.SearchUsersRequest{City: "LA"})
	if err != nil {
		log.Fatalf("failed to search users: %v", err)
	}
	log.Printf("Users found in LA: %v", searchUsersResponse.GetUsers())
}
