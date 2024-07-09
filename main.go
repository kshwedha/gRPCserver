package main

import (
	"context"
	"fmt"
	"log"
	"net"

	Otel "github.com/kshwedha/gRPCserver/otel"

	"google.golang.org/grpc"
)

type otelServer struct {
	Otel.UnimplementedOtelServer
}

func (o *otelServer) Create(ctx context.Context, req *Otel.Request) (*Otel.Response, error) {
	fmt.Printf("received a message: %s", req.String())
	return &Otel.Response{
		Reply: "this is the reply!",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("%s", err)
	}
	ServiceRegistrar := grpc.NewServer()
	service := &otelServer{}
	Otel.RegisterOtelServer(ServiceRegistrar, service)
	log.Println("Starting gRPC server on port 8089...")
	err = ServiceRegistrar.Serve(listener)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
