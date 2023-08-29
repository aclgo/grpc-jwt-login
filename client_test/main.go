package main

import (
	"context"
	"log"
	"time"

	"github.com/aclgo/grpc-jwt/proto" // Import your generated proto package
	"google.golang.org/grpc"
)

func main() {
	// Set up gRPC connection
	// creds := credentials.NewClientTLSFromCert(nil, "")
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		// grpc.WithBlock(), // Wait until connection is established
	}

	conn, err := grpc.DialContext(context.Background(), ":50051", opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := proto.NewUserServiceClient(conn)

	// Make a gRPC call
	createUserReq := &proto.CreateUserRequest{
		Email:    "dolor",
		LastName: "ut",
		Name:     "laboris",
		Password: "laborum",
		Role:     "Lorem",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Register(ctx, createUserReq, []grpc.CallOption{}...)
	if err != nil {
		log.Fatalf("Failed to register user: %v", err)
	}

	log.Printf("User registered successfully. ID: %s", resp)
}
