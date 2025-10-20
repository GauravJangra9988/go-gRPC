package main

import (
	pb "github/go-gRPC/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {

	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF{
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}

		log.Printf("Got stream: %v", req.Message)
		messages = append(messages, "Hello", req.Message)
	}


}