package main

import (
	"log"

	pb "github.com/ehernandez-xk/grpc-demo/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()

	c := pb.NewMyServiceClient(conn)

	// Add a person
	r, err := c.AddPerson(context.Background(), &pb.Person{Name: "Pepe"})
	if err != nil {
		log.Fatalf("Could not add name %v", err)
	}
	log.Printf("AddPerson: %v", r.Status)
}
