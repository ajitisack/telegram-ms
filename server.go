package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
	pb "tradebot/proto"

	"google.golang.org/grpc"
)

type TelegramServer struct {
	pb.UnimplementedTelegramMessageServiceServer
}

func (tgs *TelegramServer) SendTelegramMessage(ctx context.Context, in *pb.TelegramMessage) (*pb.TelegramMessageStatus, error) {
	t := time.Now().Format("2006-01-02 15:04:05.9999")
	message := url.QueryEscape(in.GetMessage())
	fmt.Printf("%s : Sending message [%s]\n", t, in.GetMessage())
	token := ""
	chatid := ""
	msg_url := "https://api.telegram.org/bot" + token + "/sendMessage?chat_id=" + chatid + "&parse_mode=Markdown&text=" + message
	// fmt.Printf("%s\n", msg_url)
	http.Get(msg_url)
	return &pb.TelegramMessageStatus{Done: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("failed connection: %v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterTelegramMessageServiceServer(s, &TelegramServer{})
	fmt.Printf("server listening at %v\n", lis.Addr())
	s.Serve(lis)
}
