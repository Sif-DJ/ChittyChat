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

	nodeName := "RandomNum"
	var joinReq proto.JoinRequest
	joinReq.NodeName = nodeName

	responseJoin, _ := client.Join(context.Background(), &joinReq)
	log.Println(responseJoin.Lamport, responseJoin.NodeId, responseJoin.Status)

	// I now wish to disconnect

	var leaveReq proto.LeaveRequest
	leaveReq.SenderId = nodeName

	responseLeave, _ := client.Leave(context.Background(), &leaveReq)
	log.Println(responseLeave.Lamport, responseLeave.NodeId, responseLeave.Status)
}

func logMessages(stream proto.ChittyChat_BroadcastServer) {

}
