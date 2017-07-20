package main

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"net"

	"github.com/boltdb/bolt"
	pb "github.com/ehernandez-xk/grpc-demo/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	//The server manage the session of the  database
	db *bolt.DB
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

//AddPerson Add person to the service
func (a *server) AddPerson(ctx context.Context, in *pb.Person) (*pb.Replay, error) {
	log.Println("Adding Person", in.Name)

	a.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("People"))
		if err != nil {
			return err
		}
		id, err := b.NextSequence()
		if err != nil {
			log.Fatal("Err with NextSequence", err)
		}
		intID := int(id)
		buf, err := json.Marshal(in)
		if err != nil {
			log.Fatal("Err json Marshal", err)
		}
		err = b.Put(itob(intID), buf)
		if err != nil {
			return err
		}
		return nil
	})

	return &pb.Replay{Status: "Person Added"}, nil
}

//ListPeople added in the service
func (a *server) ListPeople(ctx context.Context, in *pb.Empty) (*pb.ListReplay, error) {
	log.Println("List People")

	people := []*pb.Person{}
	a.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("People"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			//fmt.Printf("key=%s, value=%s\n", k, v)
			p := pb.Person{}
			err := json.Unmarshal(v, &p)
			if err != nil {
				return err
			}
			people = append(people, &p)
		}
		return nil
	})

	return &pb.ListReplay{People: people}, nil
}

func main() {
	log.Println("server on port: ", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//Open new session of the database
	dbSession, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer dbSession.Close()

	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{db: dbSession})
	s.Serve(lis)
}
