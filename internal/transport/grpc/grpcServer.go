package grpcServer

import (
	"net"

	"context"
	"fmt"

	p "github.com/krassor/serverHttp/internal/transport/grpc/proto/pb"
	um "github.com/krassor/serverHttp/pkg/utils"
	"google.golang.org/grpc"
)

type MessageServer struct {
}

func (MessageServer) SayIt(ctx context.Context, r *p.Request) (*p.Response, error) {
	fmt.Println("Request Text:", r.Text)
	fmt.Println("Request SubText:", r.Subtext)
	response := &p.Response{
		Text:    r.Text,
		Subtext: "Got it!",
	}
	return response, nil
}

func ServerGrpcStart(grpcPort string) error {
	um.PrintlnWithTimeShtamp("Server gRPC starting")
	portGrpc := ":" + grpcPort
	server := grpc.NewServer()
	var messageServer MessageServer
	p.RegisterMessageServiceServer(server, messageServer)
	listen, err := net.Listen("tcp", portGrpc)
	if err != nil {
		return err
	}
	um.PrintlnWithTimeShtamp("Server gRPC listening...")
	server.Serve(listen)
	defer listen.Close()
	return nil
}
