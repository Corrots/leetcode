package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	pb "github.com/corrots/leetcode/test/rpc/protos/proto/user.proto"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Received username: %s", in.Username)
	token := fmt.Sprintf("%x", md5.Sum([]byte("abc")))
	return &pb.LoginResponse{
		Username: in.Username,
		Nickname: "abc",
		Token:    token,
		Code:     http.StatusOK,
		Msg:      "succeed",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
