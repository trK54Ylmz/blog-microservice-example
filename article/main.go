package main

import (
	"log"
	"net"

	"github.com/trk54ylmz/blog-ms/proto/article"
	"google.golang.org/grpc"
)

func main() {
	conn, err := net.Listen("tcp", ":8002")

	if err != nil {
		log.Fatalf("net dial failed %s", err)

		return
	}

	// Create grpc service
	server := grpc.NewServer()

	// Create new article service
	a, err := NewArticleService()

	// Close database connection when stop
	defer a.Cancel()

	if err != nil {
		log.Fatalf("article service error %s", err)
	}

	// Register article service implementation
	article.RegisterArticleServiceServer(server, a)

	log.Println("article service starting at 8002")

	// Start grpc server
	err = server.Serve(conn)

	if err != nil {
		log.Fatalf("server start failed %s", err)

		return
	}
}
