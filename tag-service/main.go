package main

import (
	"github.com/go-programming-tour-book/tag-service/server"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "tag-service/proto"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())

	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("net.listen err:%v", err)
	}

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("server.Serve err:%v", err)
	}
}
