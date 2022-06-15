package main

import (
	"log"
	"net"

	"github.com/trk54ylmz/blog-ms/proto/user"
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

	err = server.Serve(conn)

	if err != nil {
		log.Fatalf("server start failed %s", err)
	}
}
