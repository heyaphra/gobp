package main

import (
	"context"
	"fmt"
	"log"
	"net"

	cowsay "github.com/Code-Hex/Neo-cowsay"
	"github.com/rebeccajae/gobp/internal/proto"
	"google.golang.org/grpc"
)

// 6683 is MOTD on a T9 keyboard
const port = ":6683"

type cowserver struct {
	proto.UnimplementedCowSayerServer
}

func (ms *cowserver) Cowsay(ctx context.Context, in *proto.CowsayRequest) (*proto.CowsayReply, error) {
	log.Printf("%s wants a %s", in.GetName(), in.GetCow())
	say, err := cowsay.Say(
		cowsay.Phrase(fmt.Sprintf("Hello, %s", in.GetName())),
		cowsay.Type(in.GetCow()),
		cowsay.BallonWidth(40),
	)
	if err != nil {
		return nil, err
	}
	return &proto.CowsayReply{Message: say}, nil
}

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
