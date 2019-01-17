package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/ethercflow/injuredfs/pb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", ":65534", "The address to bind to")
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewInjureClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ms, err := c.Methods(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("RPC Methods failed: %v", err)
	} else {
		log.Println("Vaild methods: ", ms)
	}
}
