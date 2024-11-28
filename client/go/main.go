package main

import (
	"fmt"
	"log"
    "flag"
    "rpcmap/client/go/rpcmap"
)

func main() {
    host := flag.String("host", "localhost:50051", "The host:port to connect to")
    flag.Parse()

	client := rpcmap.NewMapClient()
	err := client.Connect(fmt.Sprintf("%s", *host))
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer client.Close()

	err = client.Put("key1", "value1")
	if err != nil {
		log.Fatalf("Error putting value: %v", err)
	}

	err = client.Put("key2", "value2")
	if err != nil {
		log.Fatalf("Error putting value: %v", err)
	}

	value, err := client.Get("key1")
	if err != nil {
		log.Fatalf("Error getting value: %v", err)
	}
	fmt.Println("key1:", value)

	value, err = client.Get("key2")
	if err != nil {
		log.Fatalf("Error getting value: %v", err)
	}
	fmt.Println("key2:", value)

	err = client.Del("key2")
	if err != nil {
		log.Fatalf("Error deleting key: %v", err)
	}

	value, err = client.Get("key2")
	if err != nil {
		log.Fatalf("Error getting value: %v", err)
	}
	fmt.Println("key2 after delete:", value)
}

