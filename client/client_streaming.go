package main

import (
	"context"
	pb "github/go-gRPC/proto"
	"log"
)


func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NamesList){

	log.Printf("Client streaming started")

	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _,name := range names.Names{

	req := &pb.HelloRequest{
			Message: name,
		}

		err := stream.Send(req)
		if err != nil{
			log.Fatal(err)
		}

		log.Printf("Sent the request with name: %s", name )
	}

	res, err := stream.CloseAndRecv()
	log.Println("Client streaming finished")
	if err != nil {
		log.Fatalf("Error receiving: %v", err)
	}

	log.Println(res.Messages)
}

