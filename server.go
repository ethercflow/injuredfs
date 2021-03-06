package main

import (
	"context"
	"math/rand"
	"net"
	"regexp"
	"sync"
	"syscall"
	"time"

	pb "github.com/ethercflow/injuredfs/pb"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//go:generate protoc -I pb pb/injure.proto --go_out=plugins=grpc:pb

var (
	faultMap map[string]*faultContext
	fml      sync.Mutex

	methods map[string]bool
)

func init() {
	faultMap = make(map[string]*faultContext)
	initMethods()
}

type faultContext struct {
	errno  error
	random bool
	pct    uint32
	path   string
	delay  time.Duration
}

func initMethods() {
	methods = make(map[string]bool)
	methods["open"] = true
	methods["read"] = true
	methods["write"] = true
	methods["mkdir"] = true
	methods["rmdir"] = true
	methods["opendir"] = true
	methods["fsync"] = true
	methods["flush"] = true
	methods["release"] = true
	methods["truncate"] = true
	methods["getattr"] = true
	methods["chown"] = true
	methods["chmod"] = true
	methods["utimens"] = true
	methods["allocate"] = true
	methods["getlk"] = true
	methods["setlk"] = true
	methods["setlkw"] = true
	methods["statfs"] = true
	methods["readlink"] = true
	methods["symlink"] = true
	methods["create"] = true
	methods["access"] = true
	methods["link"] = true
	methods["mknod"] = true
	methods["rename"] = true
	methods["unlink"] = true
	methods["getxattr"] = true
	methods["listxattr"] = true
	methods["removexattr"] = true
	methods["setxattr"] = true
}

func randomErrno() error {
	// from E2BIG to EXFULL, notice linux only
	return syscall.Errno(rand.Intn(0x36-0x7) + 0x7)
}

func probab(percentage uint32) bool {
	return rand.Intn(99) < int(percentage)
}

func faultInject(path, method string) error {
	fml.Lock()
	fc, ok := faultMap[method]
	if !ok {
		fml.Unlock()
		return nil
	}
	fml.Unlock()

	if !probab(fc.pct) {
		return nil
	}

	if len(fc.path) > 0 {
		re, err := regexp.Compile(fc.path)
		if err != nil || !re.MatchString(path) {
			log.WithFields(log.Fields{
				"method": method,
				"fc":     fc,
				"path":   path,
			}).Warn("Invalid path")
			return nil
		}
	}

	log.WithFields(log.Fields{
		"method": method,
		"fc":     fc,
		"path":   path,
	}).Debug("Fault inject info")

	var errno error = nil
	if fc.errno != nil {
		errno = fc.errno
	} else if fc.random {
		errno = randomErrno()
	}

	if fc.delay > 0 {
		time.Sleep(fc.delay)
	}

	return errno
}

type server struct {
}

func (s *server) methods() []string {
	ms := make([]string, 0)
	for k := range methods {
		ms = append(ms, k)
	}
	return ms
}

func (s *server) Methods(ctx context.Context, in *empty.Empty) (*pb.Response, error) {
	return &pb.Response{Methods: s.methods()}, nil
}

func (s *server) RecoverAll(ctx context.Context, in *empty.Empty) (*empty.Empty, error) {
	fml.Lock()
	defer fml.Unlock()
	// The compiler(1.11) now optimizes map clearing operations of the form:
	for k := range faultMap {
		delete(faultMap, k)
	}
	return &empty.Empty{}, nil
}

func (s *server) RecoverMethod(ctx context.Context, in *pb.Request) (*empty.Empty, error) {
	ms := in.GetMethods()
	fml.Lock()
	defer fml.Unlock()
	for _, v := range ms {
		delete(faultMap, v)
	}
	return &empty.Empty{}, nil
}

func (s *server) setFault(ms []string, f *faultContext) {
	fml.Lock()
	defer fml.Unlock()
	for _, v := range ms {
		faultMap[v] = f
	}
}

func (s *server) SetFault(ctx context.Context, in *pb.Request) (*empty.Empty, error) {
	// TODO: use Errno(0), and hanle Errno(0) in Hook interfaces
	var errno error = nil
	if in.Errno != 0 {
		errno = syscall.Errno(in.Errno)
	}
	f := &faultContext{
		errno:  errno,
		random: in.Random,
		pct:    in.Pct,
		path:   in.Path,
		delay:  time.Duration(in.Delay) * time.Microsecond,
	}
	s.setFault(in.Methods, f)
	return &empty.Empty{}, nil
}

func (s *server) SetFaultAll(ctx context.Context, in *pb.Request) (*empty.Empty, error) {
	// TODO: use Errno(0), and hanle Errno(0) in Hook interfaces
	var errno error = nil
	if in.Errno != 0 {
		errno = syscall.Errno(in.Errno)
	}
	f := &faultContext{
		errno:  errno,
		random: in.Random,
		pct:    in.Pct,
		path:   in.Path,
		delay:  time.Duration(in.Delay) * time.Microsecond,
	}
	s.setFault(s.methods(), f)
	return &empty.Empty{}, nil
}

func StartServer(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterInjureServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}
