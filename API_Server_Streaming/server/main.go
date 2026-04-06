package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"server_streaming/src/pb/departament"
	"strconv"
	"strings"

	"google.golang.org/grpc"
)

type server struct {
	departament.DepartamentServiceServer
}

func (s *server) ListPerson(req *departament.ListPersonRequest, srv departament.DepartamentService_ListPersonServer) error {

	file, err := os.Open("./data.csv")
	if err != nil {
		return fmt.Errorf("Error on open file. Error: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ";")
		id, _ := strconv.Atoi(data[0])
		name := data[1]
		email := data[2]
		income, _ := strconv.Atoi(data[3])
		departamentID, _ := strconv.Atoi(data[4])

		if int32(departamentID) == req.GetDepartamentID() {
			if err := srv.Send(&departament.ListPersonResponse{
				Id:            int32(id),
				Name:          name,
				Email:         email,
				Income:        int32(income),
				DepartamentID: int32(departamentID),
			}); err != nil {
				return fmt.Errorf("Error on send. Error: %v", err)
			}
		}
	}

	return nil
}

func main() {

	fmt.Println("Starting grcp server")
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("Error. on get listener. Error: ", err)
	}

	s := grpc.NewServer()
	departament.RegisterDepartamentServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalln("Error on server. Error: ", err)
	}

}
