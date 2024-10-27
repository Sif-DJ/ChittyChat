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
	messages      []proto.Message
	lamport       proto.Lamport
	broadcastChan chan *proto.Message
}

func main() {
	srv := &ChittyChatServer{messages: []proto.Message{}}
	var message proto.Message
	message.Text = "The Server is Online!"
	srv.messages = append(srv.messages, message)

	srv.startServer()
	srv.lamport.NodeId = "Server"
	srv.lamport.Time = 0
}

func (srv *ChittyChatServer) Publish(ctx context.Context, msg *proto.Message) (*proto.PublishResponse, error) {
	srv.lamport.Time++
	srv.messages = append(srv.messages, *msg)
	log.Println(msg)
	response := proto.PublishResponse{
		Lamport: &srv.lamport,
		Status:  proto.Status_OK,
	}
	return &response, nil
}

func (srv *ChittyChatServer) Broadcast(_ *proto.BroadcastSubscription, stream proto.ChittyChat_BroadcastServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return nil
		case msg := <-srv.broadcastChan:
			if err := stream.Send(msg); err != nil {
				return err
			}
		}
	}

	/*for i := 0; i < len(srv.messages); i++ {
		stream.Send(&srv.messages[i])
	}
	stream.Send(&srv.messages[len(srv.messages)-1])
	return nil*/
}

func (srv *ChittyChatServer) Join(ctx context.Context, req *proto.JoinRequest) (*proto.JoinResponse, error) {
	log.Println(req)
	var msg proto.Message
	msg.Text = "Participant " + req.NodeName + " joined ChittyChat"
	msg.Lamport = req.Lamport
	_, err := srv.Publish(ctx, &msg)
	var response proto.JoinResponse
	response.NodeId = req.NodeName
	if err != nil {
		response.Status = proto.Status_GENERAL_ERROR
	} else {
		response.Status = proto.Status_OK
	}
	response.Lamport = &srv.lamport
	return &response, nil
}

func (srv *ChittyChatServer) Leave(ctx context.Context, req *proto.LeaveRequest) (*proto.LeaveResponse, error) {
	log.Println(req)
	var msg proto.Message
	msg.Text = "Participant " + req.SenderId + " left ChittyChat"
	msg.Lamport = req.Lamport
	_, err := srv.Publish(ctx, &msg)
	var response proto.LeaveResponse
	response.NodeId = req.SenderId
	if err != nil {
		response.Status = proto.Status_GENERAL_ERROR
	} else {
		response.Status = proto.Status_OK
	}
	response.Lamport = &srv.lamport
	return &response, nil
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
