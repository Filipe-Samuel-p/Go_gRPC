package main

import (
	"client_streaming/src/pb/calc"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	calc.CalcServiceServer
}

func (s *server) Calc(stream calc.CalcService_CalcServer) error {

	var quantity int32 = 0
	var total int32 = 0

	for {
		input, err := stream.Recv() // Este método recebe todos os inputs do cliet até ele finalizar.
		if err == io.EOF {          // Aqui eu garanto que o client terminou de enviar as coisas
			avg := float64(total / quantity)
			stream.SendAndClose(&calc.Output{ // quando o client parar de enviar, o server precisa retornar algo. Aqui fica a lógica
				Quantity: quantity,
				Averege:  avg,
				Total:    total,
			})
		}
		if err != nil {
			return nil
		}

		quantity++
		total += int32(input.GetValue())

		fmt.Print("Input: ", input)
	}

}

func main() {

	fmt.Println("Starting server gcrp")

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("Error on get listener. Error: ", err)
	}

	s := grpc.NewServer()
	calc.RegisterCalcServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalln("Error on serve. Error: ", err)
	}

}
