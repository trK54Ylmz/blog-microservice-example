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

	// Create new user service
	u, err := NewUserService()

	// Close database connection when stop
	defer u.db.Close()

	if err != nil {
		log.Fatalf("user service error %s", err)
	}

	// Register user service implementation
	user.RegisterUserServiceServer(server, u)

	log.Println("user service starting at 8001")

	// Start grpc server
	err = server.Serve(conn)

	if err != nil {
		log.Fatalf("server start failed %s", err)

		return
	}
}
