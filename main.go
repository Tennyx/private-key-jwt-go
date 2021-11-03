package main

import (
	"github.com/Tennyx/private-key-jwt-go/keygen"
	"github.com/Tennyx/private-key-jwt-go/client"
	"github.com/Tennyx/private-key-jwt-go/createjwt"
	"fmt"
)


func main() {
	fmt.Println("\nWhich task would you like to run?\n", "\n[1] Generate Keys", "\n[2] Create Okta Service App", "\n[3] Create PKJ and get Access Token\n")
	
	var input string

	fmt.Scanln(&input)
	
	if input == "1" {
		keygen.GenerateKeys()
	} else if input == "2" {
		client.RegisterOktaClient()
	} else if input == "3" {
		createjwt.GetAccessToken()
	} else {
		fmt.Println("Invalid input.")
	}
}
