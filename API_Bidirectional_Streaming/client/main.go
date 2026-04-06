package main

import (
	"bidirectional_streaming/src/pb/shoppingcart"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Error on new client. Error: ", err)
	}
	defer conn.Close()

	client := shoppingcart.NewShoppingCartServiceClient(conn) // abrindo canal com o server
	stream, err := client.AddItem(context.Background())
	if err != nil {
		log.Fatalln("Error on get channel to stream. Error: ", err)
	}

	waitch := make(chan struct{})

	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				close(waitch)
				return
			}
			if err != nil {
				log.Fatalln("Error on recv. Error: ", err)
			}

			fmt.Printf("<-- Response: %v\n", response)
		}
	}()

	items := []shoppingcart.AddProduct{
		{ProductId: 1, Quantity: 2, PriceUnit: 5.0},
		{ProductId: 2, Quantity: 5, PriceUnit: 12.0},
	}

	for _, v := range items {
		if err := stream.Send(&v); err != nil {
			log.Fatalln("Error on send response. Error: ", err)
		}

		fmt.Printf("-> Send: %v\n", v)
	}

	<-waitch

}
