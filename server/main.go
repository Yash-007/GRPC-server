package main

import (
	"context"
	"fmt"
	pb "grpc-server/user/grpc-server/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserService) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	fmt.Println("GetUser", req)
	return &pb.UserResponse{
		Name: "Yash",
		Age:  21,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &UserService{})
	fmt.Println("UserService gRPC server started on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
