package main

import (
	proto "ChittyChat/grpc"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ChittyChatServer struct {
	proto.UnimplementedChittyChatServer
	messages []proto.Message
}

func main() {
	srv := &ChittyChatServer{messages: []proto.Message{}}
	var message proto.Message
	message.Text = "The Server is Online!"
	message.Timestamp = 0
	srv.messages = append(srv.messages, message)

	srv.startServer()
}

func (srv *ChittyChatServer) Publish(ctx context.Context, msg *proto.Message) (*proto.Empty, error) {
	msg.Timestamp = int32(len(srv.messages))
	srv.messages = append(srv.messages, *msg)
	log.Println(msg)
	return new(proto.Empty), nil
}

func (srv *ChittyChatServer) GetMessages(ctx context.Context, _ *proto.Empty) ([]proto.Message, error) {
	return srv.messages, nil
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
