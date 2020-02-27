package main

import (
	"context"
	"log"
	"net"

	"github.com/joshuabezaleel/test-grpc/pb"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
	var userSrv UserServer
	pb.RegisterUsersServer(srv, userSrv)

	log.Println("Starting RPC server at", ":7000")

	l, err := net.Listen("tcp", ":7000")
	if err != nil {
		log.Fatalf("could not listen to %s: %v", ":7000", err)
	}

	log.Fatal(srv.Serve(l))
}

var localStorage *pb.UserList

func init() {
	localStorage = new(pb.UserList)
	localStorage.List = make([]*pb.User, 0)
}

type UserServer struct{}

func (UserServer) Register(ctx context.Context, user *pb.User) (*empty.Empty, error) {
	localStorage.List = append(localStorage.List, user)

	log.Println("Registering user", user.String())

	return new(empty.Empty), nil
}

func (UserServer) List(ctx context.Context, void *empty.Empty) (*pb.UserList, error) {
	return localStorage, nil
}
