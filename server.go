package main

import (
	"context"
	"log"
	"net"

	pb "github.com/uni-3/go-grpc-tutorial/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("RECV: %v", in.GetName())
	message := "Hello, " + in.GetName() + "!"
	log.Printf("SEND: %v", message)
	return &pb.HelloReply{Message: message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("serving... :5001")
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
