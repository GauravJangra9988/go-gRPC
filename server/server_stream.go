package main

import (
	pb "github/go-gRPC/proto"
	"log"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error{

	log.Printf("List of names: %v", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}

		err := stream.Send(res)
		if err != nil {
			return err
		}
	}


	return nil


}


