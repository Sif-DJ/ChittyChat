package main

import (
	proto "ITUServer/grpc"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ChittyChat struct {
	proto.UnimplementedChittyChatServer
	messages []string
}

func main() {
	server := &ChittyChat{messages: []string{}}
	server.messages = append(server.messages, "First Message")

	server.startServer()
}

func (server *ChittyChat) Publish(ctx context.Context, msg string) {
	server.messages = append(server.messages, msg)
	fmt.Println(msg)
}

func (server *ChittyChat) startServer() {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("Exception Error")
	}
	proto.RegisterChittyChatServer(grpcServer, server)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Exception Error after Registration")
	}
}
