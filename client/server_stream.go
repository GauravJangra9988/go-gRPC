package main

import (
	"context"
	pb "github/go-gRPC/proto"
	"io"
	"log"
)





func callSayHelloStream(client pb.GreetServiceClient, names *pb.NamesList){
	
	log.Printf("Streaming started")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatal(err)

	}

	for {
		message, err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		log.Println(message)
	}

	log.Println("Streaming ends")


}