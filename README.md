# gRPC User Service

## Overview

This project implements a gRPC User Service with search capabilities. It allows interaction via gRPC endpoints to manage user data.

## Prerequisites

Before you begin, ensure you have the following installed:
- Go (version >= 1.16)
- Docker (optional, for deployment)

## Installation

Clone the repository:

git clone https://github.com/yourusername/grpc-user-service.git
cd grpc-user-service

Install dependencies:
    go mod download

Running Locally
    go build -o grpc-user-service ./cmd/server
    ./grpc-user-service

The gRPC server should now be running on localhost:50051.

Access gRPC Service Endpoints
    grpcurl -plaintext localhost:50051 list

    grpcurl -plaintext -d '{"user_id": "123"}' localhost:50051 your.package.YourService/YourMethod

Configuration
    Configuration can be adjusted via environment variables or a configuration file (config.yaml).
    Ensure necessary parameters such as server port, database connection details, etc., are configured correctly.

## Deployment with Docker

### Build Docker Image
    1. docker build -t grpc-user-service .
    2. docker run -p 50051:50051 grpc-user-service
#### The gRPC server inside the Docker container will be accessible on localhost:50051.
### Configuration (Docker)
    
