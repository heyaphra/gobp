package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/user"
	"time"

	"github.com/rebeccajae/gobp/internal/proto"
	"google.golang.org/grpc"
)

const address = "localhost:6683"

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewCowSayerClient(conn)
	name := user.Name
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Cowsay(ctx, &proto.CowsayRequest{
		Name: name,
		Cow:  "stimpy",
	})
	if err != nil {
		log.Fatalf("could not moo: %v", err)
	}
	fmt.Println(r.GetMessage())

}
