package main

import (
	"flag"
	"fmt"
	"log"

	pb "github.com/ehernandez-xk/grpc-demo/service"
	"github.com/icrowley/fake"
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
	fmt.Println(address)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()

	c := pb.NewMyServiceClient(conn)

	if *option == "add" {
		// Add option adds a person
		r, err := c.AddPerson(context.Background(), &pb.Person{Name: *name})
		if err != nil {
			log.Fatalf("Could not add name %v", err)
		}
		fmt.Printf("rpc - AddPerson replay: %v\n", r.Status)
	}
	// List option lists people in the db
	if *option == "list" {
		r, err := c.ListPeople(context.Background(), &pb.Empty{})
		if err != nil {
			log.Fatalf("Could not list people")
		}
		fmt.Println("People in the database")
		for i, person := range r.People {
			fmt.Printf("%v <-- %s\n", i, person.Name)
		}
	}
	// Fake option adds 1000 fake names
	if *option == "fake" {
		for i := 0; i < 1000; i++ {
			_, err := c.AddPerson(context.Background(), &pb.Person{Name: fake.FullName()})
			if err != nil {
				log.Fatalf("Could not add name %v", err)
			}
		}
		fmt.Println("Done fake")
	}

}
