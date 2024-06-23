package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// <모듈 경로>/<패키지 경로>
	pb "grpc-test/grpc-test/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	http.HandleFunc("/test", func(write http.ResponseWriter, rsponse *http.Request) {

		server := os.Getenv("SERVER")
		port := os.Getenv("PORT")

		if server == "" {
			server = "localhost"
		}
		if port == "" {
			port = "50051"
		}

		fmt.Println("server: ", server)
		fmt.Println("port: ", port)
		// Set up a connection to the server.
		conn, err := grpc.NewClient(server+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewGreeterClient(conn)

		// Contact the server and print out its response.
		name := "world"
		if len(os.Args) > 1 {
			name = os.Args[1]
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
		fmt.Fprintf(write, "Greeting: %s", r.GetMessage())
	})

	log.Println("HTTP server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
