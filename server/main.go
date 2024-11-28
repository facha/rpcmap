package main

import (
    "fmt"
	"log"
	"flag"
	"net"
	"google.golang.org/grpc"
    "rpcmap/server/mapservice"
)

func main() {
	port := flag.String("port", "50051", "The port to bind the server to")
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", *port, err)
	}

	grpcServer := grpc.NewServer()
	mapService := NewMapService() 
	rpcServer := NewRPCServer(mapService)
	mapservice.RegisterMapServiceServer(grpcServer, rpcServer)

	log.Printf("Server is running on port :%s", *port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

