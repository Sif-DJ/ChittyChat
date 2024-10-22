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
	srv := &ChittyChatServer{messages: []string{}}
	srv.messages = append(srv.messages, "First Message")

	srv.startServer()
}

func (srv *ChittyChatServer) Publish(ctx context.Context, msg string) {
	srv.messages = append(srv.messages, msg)
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
