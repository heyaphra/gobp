package main

import (
	"context"
	"fmt"
	"testing"

	cowsay "github.com/Code-Hex/Neo-cowsay"
	"github.com/rebeccajae/gobp/internal/proto"
)

func TestCowsay(t *testing.T) {
	cowsayRequest := &proto.CowsayRequest{
		Name: "Rebecca",
		Cow:  "default",
	}

	// Trust that the imported library's tests are robust
	expectedOutput, err := cowsay.Say(
		cowsay.Phrase(fmt.Sprintf("Hello, %s", cowsayRequest.GetName())),
		cowsay.Type(cowsayRequest.GetCow()),
		cowsay.BallonWidth(40),
	)
	if err != nil {
		t.Error(err)
	}

	cowsayServer := &cowserver{}
	ret, err := cowsayServer.Cowsay(context.Background(), cowsayRequest)
	if err != nil {
		t.Error(err)
	}

	if ret.GetMessage() != expectedOutput {
		t.Errorf("Expected:\n%s\nGot:\n%s\n", expectedOutput, ret.GetMessage())
	}
}
