package main

import (
	proto "ChittyChat/grpc"
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Exception Error")
	}

	client := proto.NewChittyChatClient(conn)

	randomNum := rand.Intn(265)
	nodeName := fmt.Sprint(randomNum)
	var lamport proto.Lamport
	lamport.Time = 1
	lamport.NodeId = nodeName

	var joinReq proto.JoinRequest
	joinReq.NodeName = nodeName
	joinReq.Lamport = &lamport

	responseJoin, _ := client.Join(context.Background(), &joinReq)
	lamport.Time = responseJoin.Lamport.Time
	log.Println(responseJoin.Lamport, responseJoin.NodeId, responseJoin.Status)

	// Being logging
	stream, _ := client.Broadcast(context.Background(), &proto.BroadcastSubscription{
		Receiver: nodeName,
	})
	go logMessages(stream)

	// I wish to text
	for {
		var input string
		reader := bufio.NewReader(os.Stdin)
		input, _ = reader.ReadString('\n')
		before, _, _ := strings.Cut(input, "\r\n")
		if before == "exit" {
			break
		}
		if len(before) <= 128 {
			publishMessage(client, &lamport, before)
		} else {
			fmt.Println("Error: Too many characters in message to send! Max is 128 characters")
		}
	}

	// I now wish to disconnect

	var leaveReq proto.LeaveRequest
	leaveReq.SenderId = nodeName
	leaveReq.Lamport = &lamport

	responseLeave, _ := client.Leave(context.Background(), &leaveReq)
	log.Println(responseLeave.Lamport, responseLeave.NodeId, responseLeave.Status)
}

func logMessages(stream proto.ChittyChat_BroadcastClient) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Println(msg)
	}
}

func publishMessage(client proto.ChittyChatClient, lamport *proto.Lamport, text string) {

	msg := &proto.Message{
		Text:    text,
		Lamport: lamport,
	}
	response, _ := client.Publish(context.Background(), msg)

	lamport.Time = response.Lamport.Time
	log.Println("Message published. Status: " + response.Status.String())
}
