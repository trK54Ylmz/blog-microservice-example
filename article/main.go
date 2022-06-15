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

	server := grpc.NewServer()

	article.RegisterArticleServiceServer(server, new(ArticleService))

	log.Println("article service starting at 8082")

	// Start grpc server
	err = server.Serve(conn)

	if err != nil {
		log.Fatalf("server start failed %s", err)

		return
	}
}
