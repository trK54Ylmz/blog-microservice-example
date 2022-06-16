package main

import (
	"log"
	"net"

	"github.com/trk54ylmz/blog-microservice-example/proto/user"
	"google.golang.org/grpc"
)

func main() {
	conn, err := net.Listen("tcp", ":8001")

	if err != nil {
		log.Fatalf("net dial failed %s", err)

		return
	}

	server := grpc.NewServer()

	user.RegisterUserServiceServer(server, new(UserService))

	log.Println("user service starting at 8001")

	// Start grpc server
	err = server.Serve(conn)

	if err != nil {
		log.Fatalf("server start failed %s", err)

		return
	}
}
