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

type uploadServer struct {
	pb.UploadServiceServer
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}



	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	pb.RegisterUploadServiceServer(grpcServer, &uploadServer{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server Started")

}
