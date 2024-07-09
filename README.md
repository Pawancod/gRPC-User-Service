# gRPC User Service

## Overview

This project implements a gRPC User Service with search capabilities. It allows interaction via gRPC endpoints to manage user data.

## Prerequisites

Before you begin, ensure you have the following installed:
- Go (version >= 1.16)
- Docker (optional, for deployment)

## Installation

### Clone the repository:

    1.git clone https://github.com/yourusername/grpc-user-service.git
    2.cd grpc-user-service

### Install dependencies:
    1.go mod download

### Running Locally
    1.go build -o grpc-user-service ./cmd/server
    2./grpc-user-service

#### The gRPC server should now be running on localhost:50051.

### Access gRPC Service Endpoints
    1.grpcurl -plaintext localhost:50051 list
    2.grpcurl -plaintext -d '{"user_id": "123"}' localhost:50051 your.package.YourService/YourMethod

#### Configuration
    1.Configuration can be adjusted via environment variables or a configuration file (config.yaml).(optional-can be dobe later)
    2.Ensure necessary parameters such as server port, database connection details, etc., are configured correctly.

## Deployment with Docker

### Build Docker Image
    1. docker build -t grpc-user-service .
    2. docker run -p 50051:50051 grpc-user-service
#### The gRPC server inside the Docker container will be accessible on localhost:50051.


