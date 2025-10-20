package main


import (
	"context"
	pb "github/go-gRPC/proto"
	"io"
	"log"
)

func callSayHelloBidirectional(client pb.GreetServiceClient, names *pb.NamesList){

	log.Println("Streaming to server started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	
	for _, name := range names.Names{
		req := &pb.HelloRequest{
			Message: name,
		}
		err := stream.Send(req)
		if err != nil {
			log.Fatal(err)
		}
		
	}

	log.Println("Stream data to server successfull")
	log.Println("Now listening stream from server")

	stream.CloseSend()

	for {
		messge, err := stream.Recv()
		if err == io.EOF{
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Stream message received: %s", messge.Message)
	}

	log.Println("Streaming ends")
}