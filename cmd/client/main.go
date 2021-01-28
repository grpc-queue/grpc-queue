package main

import (
	"fmt"
	"os"

	"github.com/grpc-queue/grpc-queue/internal/client"
)

func main() {
	c := client.NewClient(":9000")
	fmt.Println("Server response: " + c.Send(os.Args[1]).Body)

}
