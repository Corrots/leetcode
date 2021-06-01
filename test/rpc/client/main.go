package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/corrots/leetcode/test/rpc/protos/proto/user.proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := pb.NewUserServiceClient(conn)
	parameter := &pb.LoginRequest{
		Username: "abc",
		Password: "test",
	}
	resp, err := c.Login(ctx, parameter)
	if err != nil {
		log.Fatalf("login failed: %v", err)
	}
	log.Printf("Username: %s  Code: %d   Msg: %s", resp.GetUsername(), resp.Code, resp.Msg)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
