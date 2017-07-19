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

//AddPerson Add person to the service
func (a *server) AddPerson(ctx context.Context, in *pb.Person) (*pb.Replay, error) {
	log.Println("Adding Person", in.Name)
	return &pb.Replay{Status: "Person Added"}, nil
}

//ListPeople added in the service
func (a *server) ListPeople(ctx context.Context, in *pb.Empty) (*pb.ListReplay, error) {
	log.Println("List People")

	//Temporal Person, this will come from the database
	p := pb.Person{
		Name: "Eddy",
	}
	p2 := pb.Person{
		Name: "pepe",
	}
	persons := []*pb.Person{}
	persons = append(persons, &p, &p2)

	return &pb.ListReplay{People: persons}, nil
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
