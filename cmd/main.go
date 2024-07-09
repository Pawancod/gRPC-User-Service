package main

import "github.com/Pawancod/gRPC-User-Service/internal/server"

func main() {
    server.StartGRPCServer(":50051")
}
