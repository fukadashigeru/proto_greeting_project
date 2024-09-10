// server/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "proto_greeting_project/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetingServiceServer
}

func (s *server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.GreetResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetingServiceServer(s, &server{})
	fmt.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
