package main

import (
	"context"
	"flag"
	"log"
	"os"
	"syscall"
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
	flag.Parse()

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

	dir := "latency/"
	if err := os.Mkdir(*mountpoint + dir, 0755); err != nil && err != syscall.EEXIST {
		log.Fatalf("Mkdir failed: %v", err)
	}
	f, err := os.Create(*mountpoint + dir + "tf")
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

	req = &pb.Request{Methods:[]string{"read"}, Errno: 0x5, Random: false, Pct: 100, Path: dir, Delay: 0}
	if _, err := c.SetFault(ctx, req); err != nil {
		log.Fatalf("Set fault failed: %v", err)
	}
	if _, err := f.Read(make([]byte, 4096)); err != nil && err == syscall.EIO {
		log.Println("Set fault syscall.EIO successfully")
	} else {
		log.Fatal("Set fault syscall.EIO failed, please check")
	}
}