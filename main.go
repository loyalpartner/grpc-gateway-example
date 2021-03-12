package main

import (
	"context"
	"fmt"

	hellowordpb "github.com/loyalpartner/grpc-gateway-example/proto/helloworld"
)

type server struct{}

func NewServer() *server {
	return &server{}
}

// func (s *server) SayHello(ctx context.Context, in *hello)
func (s *server) SayHello(cxt context.Context, in *hellowordpb.HelloRequest) (*helloworldpb.HelloReply, error) {

}

func main() {

}
