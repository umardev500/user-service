package main

import (
	"customer/injector"
	"customer/pb"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// get the port number from the environment
	port := os.Getenv("PORT")

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	gRPC := grpc.NewServer()
	reflection.Register(gRPC)

	handler := injector.NewUserInjector()
	pb.RegisterUserServiceServer(gRPC, handler)

	fmt.Printf("⚡️[server]: gRPC Server is running on port %s\n", port)
	log.Fatal(gRPC.Serve(listener))
}
