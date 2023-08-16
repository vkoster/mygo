package main

import (
	"google.golang.org/grpc"
	"log"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr)

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	// Connection schließen, bevor das Programm beendet
	defer conn.Close()
	//..
}
