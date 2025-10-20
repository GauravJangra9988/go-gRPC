package main

import (
	pb "github/go-gRPC/proto"
	"io"
	"log"
)


func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {

	log.Println("Listening stream started")

	var messages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
		log.Printf("Got stream data: %s", req.Message)
		messages = append(messages, req.Message)
	}

	log.Println("Sending back stream")
	
	for _, name := range messages {
		res := &pb.HelloResponse{
			Message: name,
		}

		err := stream.Send(res)
		if err != nil {
			return err
		}
	}

	log.Println("Sending stream end")

	return nil
}