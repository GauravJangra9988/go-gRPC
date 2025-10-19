package main

import (
	pb "github/go-gRPC/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	err = grpcServer.Serve(lis)
	if err != nil{
		log.Fatal(err)
	}

}	