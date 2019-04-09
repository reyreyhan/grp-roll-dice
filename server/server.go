package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	pb "roll-dice/proto"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) RollDice(ctx context.Context, in *pb.ClientRequest) (*pb.ServerResponse, error) {
	clientDice, _ := strconv.Atoi(in.Dice)

	rand.Seed(time.Now().Unix() + time.Now().Unix())
	serverDice := rand.Intn(6-1) + 1

	message := "Client Dice : " + strconv.Itoa(clientDice) + " Your Dice : " + strconv.Itoa(serverDice)

	if clientDice < serverDice {
		fmt.Println("Client Dice : ", clientDice, " Your Dice : ", serverDice, " You Win!")
		return &pb.ServerResponse{Message: message + " You Lose!"}, nil
	} else if clientDice > serverDice {
		fmt.Println("Client Dice : ", clientDice, " Your Dice : ", serverDice, " You Lose!")
		return &pb.ServerResponse{Message: message + " You Win!"}, nil
	} else {
		fmt.Println("Client Dice : ", clientDice, " Your Dice : ", serverDice, " Draw!")
		return &pb.ServerResponse{Message: message + " Draw!"}, nil
	}

}

func main() {
	port := ":8888"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
