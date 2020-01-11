package main

import (
	"log"
	"net"

	"github.com/rebeccajae/gobp/internal/proto"
	"google.golang.org/grpc"
)

// 6683 is MOTD on a T9 keyboard
const port = ":6683"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterCowSayerServer(s, &cowserver{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
