package main

import (
	"fmt"
	"os"
)

const usage = "Usage : [username] [password]"
const useErr = "You need to enter your first name %q"

func main() {
	args := os.Args

	//	var user := Jason
	//	var password := Password

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}
	u, p := args[1], args[2]
	if u != "Jason" {
		fmt.Printf(useErr, u)
	} else if p != "Password" {
		fmt.Println("Password does not match")
	} else {
		fmt.Printf("You have entered the matrix %q %q", u, p)

	}

}
