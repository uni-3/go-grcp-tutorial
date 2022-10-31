package main

import (
	"context"
	"log"
	"time"

	pb "github.com/uni-3/go-grpc-tutorial/proto"
	"google.golang.org/grpc"
)

func main() {
	// Set up a gRPC client
	log.Printf("connect")
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Request to gRPC server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "yeah"})
	if err != nil {
		log.Fatalf("could not find user: %v", err)
	}
	log.Printf("User: %v", r.Message)
}
