# InjuredFS

A fuse based fault injection filesystem with a gRPC interface for instrumentation.


# Prerequisites

For centos 7， run `./prepare.sh`.


# Building

`make clean && make`

# Using

Load fuse module, if it is not loaded
```sh
modprobe fuse
```
Create mount directory for injuredfs.       
Notice: Faults can be applied for files that are manipulated through this directory only!  
Eg:
```sh
mkdir /fuse
```
`original` directory on the file system where actual files will be stored
Running injuredfs
```sh
./injuredfs -original /original -mountpoint /fuse
```

# interface
``` go
type InjureClient interface {
        Methods(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Response, error)
        RecoverAll(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
        RecoverMethod(ctx context.Context, in *Request, opts ...grpc.CallOption) (*empty.Empty, error)
        SetFault(ctx context.Context, in *Request, opts ...grpc.CallOption) (*empty.Empty, error)
        SetFaultAll(ctx context.Context, in *Request, opts ...grpc.CallOption) (*empty.Empty, error)
}

type Request struct {
        Methods              []string `protobuf:"bytes,1,rep,name=methods,proto3" json:"methods,omitempty"`
        Errno                uint32   `protobuf:"varint,2,opt,name=errno,proto3" json:"errno,omitempty"`
        Random               bool     `protobuf:"varint,3,opt,name=random,proto3" json:"random,omitempty"`
        Pct                  uint32   `protobuf:"varint,4,opt,name=pct,proto3" json:"pct,omitempty"`
        Path                 string   `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
        Delay                uint32   `protobuf:"varint,6,opt,name=delay,proto3" json:"delay,omitempty"`
        XXX_NoUnkeyedLiteral struct{} `json:"-"`
        XXX_unrecognized     []byte   `json:"-"`
        XXX_sizecache        int32    `json:"-"`
}

type Response struct {
        Methods              []string `protobuf:"bytes,1,rep,name=methods,proto3" json:"methods,omitempty"`
        XXX_NoUnkeyedLiteral struct{} `json:"-"`
        XXX_unrecognized     []byte   `json:"-"`
        XXX_sizecache        int32    `json:"-"`
}
```
Here's an example [main.go](examples/main.go)
