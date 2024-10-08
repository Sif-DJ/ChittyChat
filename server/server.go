package main

import (
	proto "ChittyChat/grpc"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ChittyChatServer struct {
	proto.UnimplementedChittyChatServer
	messages []string
}

func main() {
	server := &ChittyChatServer{messages: []string{}}
	server.messages = append(server.messages, "First Message")

	server.startServer()
}

func (server *ChittyChatServer) Publish(ctx context.Context, msg string) {
	server.messages = append(server.messages, msg)
	fmt.Println(msg)
}

func (srv *ChittyChatServer) startServer() {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("Exception Error")
	}
	proto.RegisterChittyChatServer(grpcServer, srv)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Exception Error after Registration")
	}
}
