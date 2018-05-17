package main

import (
    "net"
    "fmt"

    pb "helloworld" 

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
)

const (
    
    Address = "localhost:50051"
)


type greeterService struct{}


var GreeterService = greeterService{}

func (h greeterService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    resp := new(pb.HelloReply)
    resp.Message = "Hello " + in.Name + "."
    fmt.Printf("Ok")
    return resp, nil
}

func main() {
    listen, err := net.Listen("tcp", Address)
    if err != nil {
        grpclog.Fatalf("failed to listen: %v", err)
    }

    
    s := grpc.NewServer()

    
    pb.RegisterGreeterServer(s, GreeterService)

    grpclog.Println("Listen on " + Address)

    s.Serve(listen)
}
