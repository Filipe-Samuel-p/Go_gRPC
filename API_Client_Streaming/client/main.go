package main

import (
	"client_streaming/src/pb/calc"
	"context"
	"fmt"
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

	client := calc.NewCalcServiceClient(conn) // abrindo canal com o server
	stream, err := client.Calc(context.Background())
	if err != nil {
		log.Fatalln("Error on get channel to stream. Error: ", err)
	}

	nums := []int32{2, 3, 4, 1}

	for _, v := range nums {
		if err := stream.Send(&calc.Input{Value: v}); err != nil { // enviando as mensagens pro server
			log.Fatalln("Error in send. Error: ", err)
		}
	}

	response, err := stream.CloseAndRecv() // finalizando a conexão
	if err != nil {
		log.Fatalln("Error on close stream. Error: ", err)
	}

	fmt.Print("Response: ", response)

}
