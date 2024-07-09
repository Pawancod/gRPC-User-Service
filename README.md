# gRPC User Service

## Overview

This project implements a gRPC User Service with search capabilities. It allows interaction via gRPC endpoints to manage user data.

## Prerequisites

Before you begin, ensure you have the following installed:
- Go (version >= 1.16)
- Docker (optional, for deployment)

## Installation

### Clone the repository:

    1.git clone https://github.com/Pawancod/gRPC-User-Service.git
    2.cd gRPC-User-Service

### Install dependencies:
    1.go mod download

### Running Locally
    1.go build -o gRPC-User-Service ./cmd/main.go
    2./grpc-user-service

#### The gRPC server should now be running on localhost:50051.

### Access gRPC Service Endpoints
    1.grpcurl -plaintext localhost:50051 list
    2.grpcurl -plaintext -d '{"user_id": "123"}' localhost:50051 your.package.YourService/YourMethod
### OR(if grpccurl didnt work) 
    1.go run main.go(prompt for query, for examples for getting by id or search users by city)

#### Configuration
    1.Configuration can be adjusted via environment variables or a configuration file (config.yaml).(optional-can be dobe later)
    2.Ensure necessary parameters such as server port, database connection details, etc., are configured correctly.

## Deployment with Docker

### Build Docker Image
    1. docker build -t gRPC-User-Service.
    2. docker run -p 50051:50051 gRPC-User-Service.
#### The gRPC server inside the Docker container will be accessible on localhost:50051.


