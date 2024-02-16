package cmd

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/GoSeoTaxi/olegMicroserviceChatServer/grpc/pkg/chat_v1"
	"github.com/GoSeoTaxi/olegMicroserviceChatServer/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	desc.UnimplementedChatV1Server
}

// RunService запускает сервис Auth
func RunService() {

	config := config.NewConfig()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {

	log.Printf("%+v \n", req)

	return &emptypb.Empty{}, nil
}

/*
proto test
{
  "from": "John",
  "text": "Hello, world!",
  "timestamp": {
    "seconds": 1645959200,
    "nanos": 999
  }
}

*/
