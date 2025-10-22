package main

import (
	// pb "github/go-gRPC/proto"
	"log"
	pb "github/go-gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {

	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal()
	}

	defer conn.Close()


	// clientHelloResponse := pb.NewGreetServiceClient(conn)
	clientFileUpload := pb.NewUploadServiceClient(conn)

	// names := &pb.NamesList{
	// 	Names: []string{"A","B","C"},
	// }

	// callSayHello(clientHelloResponse)

	// callSayHelloStream(clientHelloResponse, names)

	// callSayHelloClientStreaming(clientHelloResponse, names)

	// callSayHelloBidirectional(clientHelloResponse, names)

	callUploadFile(clientFileUpload, "./client.txt")
}