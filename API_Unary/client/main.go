package main

import (
	"API_Unary/src/pb/products"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Error on get client. Error: ", err)
	}
	defer conn.Close()

	findAllProducts(conn)

}

func findAllProducts(conn *grpc.ClientConn) {
	productClient := products.NewProductServiceClient(conn)

	productList, err := productClient.FindAll(context.Background(), &products.Product{})
	if err != nil {
		log.Fatalln("Error on list products. Error: ", err)
	}

	fmt.Print("Products: ", productList)
}
