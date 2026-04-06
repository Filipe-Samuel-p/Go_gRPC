package main

import (
	
	"golang_grcp/src/pb/users"
	"fmt"
)

func main () {
	user1 := users.User{
		Id: 1,
		Name: "Filipe",
		Email: "felipe@gmail.com",
		Password: "1234",
	}

	fmt.Print("User: %v\n", user1)
}