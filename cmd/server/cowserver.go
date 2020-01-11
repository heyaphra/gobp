package main

import (
	"context"
	"fmt"
	"log"

	cowsay "github.com/Code-Hex/Neo-cowsay"
	"github.com/rebeccajae/gobp/internal/proto"
)

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
