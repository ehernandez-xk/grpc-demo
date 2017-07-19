package main

import (
	"log"
	"net"

	pb "github.com/ehernandez-xk/grpc-demo/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (a *server) AddPerson(ctx context.Context, in *pb.Person) (*pb.Replay, error) {
	log.Println("Adding Person", in.Name)
	return &pb.Replay{Status: "Added"}, nil
}

func main() {
	log.Println("server on port: ", port)
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{})
	s.Serve(lis)
}
