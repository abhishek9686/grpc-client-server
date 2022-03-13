package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/abhishek9686/grpc-client-server/user"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8008, "The server port")
)

func main() {
	flag.Parse()
	fmt.Println("PORT:", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := user.Server{}
	grpcServer := grpc.NewServer()
	user.RegisterUserDetailsServer(grpcServer, &s)
	fmt.Printf("Starting to serve on %s\n", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
