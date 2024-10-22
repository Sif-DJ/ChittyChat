package main

import (
	proto "ChittyChat/grpc"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Exception Error")
	}

	client := proto.NewChittyChatClient(conn)

	message := new(proto.Message)
	message.Text = "Participant [Client_Test] joined ChittyChat"
	message.Timestamp = 1

	client.Publish(context.Background(), message)
}
