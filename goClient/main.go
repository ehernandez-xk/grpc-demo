package main

import (
	"flag"
	"log"

	pb "github.com/ehernandez-xk/grpc-demo/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = "50051"
)

func main() {

	//flags
	target := flag.String("target", "localhost", "target to connect")
	name := flag.String("name", "no-name", "Your name")
	option := flag.String("option", "list", "Option of the rpc (list of add)")
	flag.Parse()

	address := *target + ":" + port
	log.Println(address)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()

	c := pb.NewMyServiceClient(conn)

	if *option == "add" {
		// Add a person
		r, err := c.AddPerson(context.Background(), &pb.Person{Name: *name})
		if err != nil {
			log.Fatalf("Could not add name %v", err)
		}
		log.Printf("rpc - AddPerson replay: %v", r.Status)
	}
	if *option == "list" {
		r, err := c.ListPeople(context.Background(), &pb.Empty{})
		if err != nil {
			log.Fatalf("Could not list persons")
		}
		log.Printf("grpc - ListPerson replay %v", r.People)
	}

}
