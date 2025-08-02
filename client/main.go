package main

import (
	"context"
	"fmt"
	"time"

	pb "grpc-server/user/grpc-server/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Failed to connect to server: ", err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	cxt, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetUser(cxt, &pb.UserRequest{Id: "404"})
	if err != nil {
		fmt.Println("Failed calling get user: ", err)
	}

	fmt.Printf("User: %+v\n", res)

}
