package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/joshuabezaleel/test-grpc/pb"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func serviceUser() pb.UsersClient {
	port := ":7000"
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return pb.NewUsersClient(conn)
}

func main() {
	user1 := pb.User{
		Id:       "n001",
		Name:     "Noval Agung",
		Password: "kw8d hl12/3m,a",
		Gender:   pb.UserGender_FEMALE,
	}

	userClient := serviceUser()

	fmt.Println("\n", "===========> user test")

	// register user1
	userClient.Register(context.Background(), &user1)
	// user.Register(context.Background(), &user2)

	res1, err := userClient.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}
	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))
}
