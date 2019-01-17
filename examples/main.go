package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	pb "github.com/ethercflow/injuredfs/pb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", ":65534", "The address to bind to")
	mountpoint = flag.String("mountpoint", "", "MOUNTPOINT")
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
		log.Println("Vaild methods: ", ms.Methods)
	}

	dir := *mountpoint + "/latency/"
	if err := os.Mkdir(dir, 0755); err != nil {
		log.Fatalf("Mkdir failed: %v", err)
	}
	f, err := os.Create(dir + "testfile")
	if err != nil {
		log.Fatalf("Create file failed: %v", err)
	}

	req := &pb.Request{Methods:[]string{"write"}, Errno: 0, Random: false, Pct: 100, Path: dir, Delay: 1000000}
	if _, err := c.SetFault(ctx, req); err != nil {
		log.Fatalf("Set fault failed: %v", err)
	}
	begin := time.Now()
	if _, err := f.WriteString("hello world"); err != nil {
		log.Fatalf("Write file failed: %v", err)
	}
	end := time.Now()
	log.Println("Write latency: ", end.Sub(begin))
}