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
	messages []proto.Message
}

type testing struct {
	message string
	time    int
}

func main() {
	srv := &ChittyChatServer{messages: []proto.Message{}}
	var message proto.Message
	message.Text = "this is a test"
	message.Timestamp = 1
	srv.messages = append(srv.messages, message)

	srv.startServer()
}

func (srv *ChittyChatServer) Publish(ctx context.Context, msg *proto.Message) (*proto.Empty, error) {
	srv.messages = append(srv.messages, *msg)
	fmt.Println(msg)
	return new(proto.Empty), nil
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
