package main

import (
	"fmt"
	"github.com/TsuchiyaYugo/grpc-example-go/pb"
	"github.com/TsuchiyaYugo/grpc-example-go/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	fmt.Println("start server!")
	port := 8080
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	server := grpc.NewServer()
	pb.RegisterRockPaperScissorsServiceServer(server, service.NewRockPaperScissorsService())

	reflection.Register(server)

	fmt.Println("listening on port:8080...")
	if err := server.Serve(listenPort); err != nil {
		return
	}
}
