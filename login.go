package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	//	var user := Jason
	//	var password := Password

	if len(args) != 3 {
		fmt.Println("You need to enter a user and a password")
		return
	}
	u, p := args[1], args[2]
	if u != "Jason" {
		fmt.Printf("user %q not found in database. ", u)
	} else if p != "Password" {
		fmt.Println("Password does not match")
	} else {
		fmt.Printf("You have entered the matrix %q %q", u, p)

	}

}
