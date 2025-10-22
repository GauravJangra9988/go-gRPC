package main

import (
	"context"
	pb "github/go-gRPC/proto"
	"io"
	"log"
	"os"
)


func callUploadFile(client pb.UploadServiceClient, filePath string) error {

	log.Printf("Uploading file: %s", filePath)

	stream, err := client.UploadFile(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(filePath)
	if err != nil{
		log.Fatal(err)
		return err
	}

	defer file.Close()


	buf := make([]byte, 5)
	batchNumber := 1

	for {
		n, err := file.Read(buf)
		if err == io.EOF{
			break
		}
		if err != nil {
			return err
		}

		chunk := buf[:n]

		err = stream.Send(&pb.UploadStream{Chunks: chunk, FilePath: filePath})
		if err != nil {
			return err
		}

		log.Printf("Sent batch #%v - Size - %v\n", batchNumber, len(chunk))
		batchNumber += 1
	}

		log.Println("File uploading complete")

		res, err := stream.CloseAndRecv()
		if err != nil {
			return err
		}

		log.Println(res.Message)

		return nil

}