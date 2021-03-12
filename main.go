package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	helloworldpb "github.com/loyalpartner/grpc-gateway-example/proto/helloworld"
)

type server struct {
	// https://github.com/grpc/grpc-go/issues/3794
	helloworldpb.UnimplementedGreeterServer
}

// func NewServer() *server {
// 	return &server{}
// }

func (s *server) SayHello(cxt context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{Message: in.Name + " world"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalln("Failed to listen", err)
	}

	s := grpc.NewServer()
	helloworldpb.RegisterGreeterServer(s, &server{})

	log.Println("Serving gRPC on 0.0.0.0:9999")
	go func(){
		log.Fatal(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(context.Background(), "0.0.0.0:9999", grpc.WithBlock(), grpc.WithInsecure())

	if err != nil {
		log.Fatal("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	err = helloworldpb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr: ":8999",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8999")
	log.Fatalln(gwServer.ListenAndServe())
}
