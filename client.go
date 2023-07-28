package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	pb "tradebot/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}

	defer conn.Close()

	c := pb.NewTelegramMessageServiceClient(conn)
	ctx := context.Background()

	for {
		fmt.Printf("\nMessage:")
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		message = strings.TrimSuffix(message, "\n")
		if message == "exit" {
			break
		}
		t := time.Now().Format("2006-01-02 15:04:05.9999")
		fmt.Printf("%s : Sending message [%s]\n", t, message)
		res, _ := c.SendTelegramMessage(ctx, &pb.TelegramMessage{User: "ajit", Message: message})
		t = time.Now().Format("2006-01-02 15:04:05.9999")
		fmt.Printf("%s : Message Status : %v\n", t, res.Done)
	}
}
