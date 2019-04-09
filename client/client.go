package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	pb "roll-dice/proto"

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	// address := "192.168.0.7:8888"
	address := "localhost:8888"
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewDiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("Start Roll Your Dice - Good Luck!")
	rand.Seed(time.Now().Unix())
	dice := strconv.Itoa(rand.Intn(6-1) + 1)
	r, err := c.RollDice(ctx, &pb.ClientRequest{Dice: dice})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Println("Server Sand Back : ", r.Message)
	// fmt.Scanln()
}
