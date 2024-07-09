package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Pawancod/gRPC-User-Service/proto/github.com/Pawancod/gRPC-User-Service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client instance
	client := proto.NewUserServiceClient(conn)

	// Create a reader for user input
	reader := bufio.NewReader(os.Stdin)

	for {
		// Display menu options
		fmt.Println("\nSelect an operation:")
		fmt.Println("1: GetUser")
		fmt.Println("2: GetUsers")
		fmt.Println("3: SearchUsers")
		fmt.Println("4: Exit")
		fmt.Print("Enter choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter user ID: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, err := strconv.Atoi(idStr)
			if err != nil {
				log.Printf("Invalid user ID: %v", err)
				continue
			}

			// Call GetUser
			getUserResponse, err := client.GetUser(context.Background(), &proto.GetUserRequest{Id: int32(id)})
			if err != nil {
				log.Printf("Failed to get user: %v", err)
				continue
			}
			log.Printf("User: %v", getUserResponse.GetUser())

		case "2":
			fmt.Print("Enter user IDs (comma-separated): ")
			idsStr, _ := reader.ReadString('\n')
			idsStr = strings.TrimSpace(idsStr)
			idStrs := strings.Split(idsStr, ",")
			var ids []int32
			for _, idStr := range idStrs {
				id, err := strconv.Atoi(strings.TrimSpace(idStr))
				if err != nil {
					log.Printf("Invalid user ID: %v", err)
					continue
				}
				ids = append(ids, int32(id))
			}

			// Call GetUsers
			getUsersResponse, err := client.GetUsers(context.Background(), &proto.GetUsersRequest{Ids: ids})
			if err != nil {
				log.Printf("Failed to get users: %v", err)
				continue
			}
			log.Printf("Users: %v", getUsersResponse.GetUsers())

		case "3":
			fmt.Print("Enter city: ")
			city, _ := reader.ReadString('\n')
			city = strings.TrimSpace(city)

			// Call SearchUsers
			searchUsersResponse, err := client.SearchUsers(context.Background(), &proto.SearchUsersRequest{City: city})
			if err != nil {
				log.Printf("Failed to search users: %v", err)
				continue
			}
			log.Printf("Users found in %s: %v", city, searchUsersResponse.GetUsers())

		case "4":
			fmt.Println("Exiting program.")
			return

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}