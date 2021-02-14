package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"

// 	"github.com/grpc-queue/grpc-queue/pkg/client"
// )

// func main() {
// 	c := client.NewClient(":9000")
// 	if len(os.Args) == 1 {
// 		panic("usage: main.go op ...params")
// 	}
// 	operation := os.Args[1]
// 	switch operation {
// 	case "hello":
// 		fmt.Println(c.Send(os.Args[2]).Body)
// 	case "create":
// 		{
// 			i, err := strconv.Atoi(os.Args[3])
// 			if err != nil {
// 				panic("invalid partition number")
// 			}
// 			fmt.Println(c.CreateStream(os.Args[2], int32(i)))
// 		}
// 	default:
// 		panic("op not implemented")

// 	}

// }
