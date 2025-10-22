package main

import (
	pb "github/go-gRPC/proto"
	"io"
	"log"
	"os"
)

func (s *uploadServer) UploadFile(stream pb.UploadService_UploadFileServer) error {

	var fileBytes []byte

	for{
		chunk,err := stream.Recv()
		if err == io.EOF{
			break
		}

		if err != nil {
			return err
		}

		fileBytes = append(fileBytes, chunk.GetChunks()...)
	}

	log.Println("Receiving file from client")


	f, err := os.Create("./client.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.Write(fileBytes)
	if err != nil {
		return err
	}

	log.Println("File Received Successfully")

	return stream.SendAndClose(&pb.UploadResponse{Message: "File Uploaded"})
}
