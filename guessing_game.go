package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	var userInput int
	var secretNumber int
	secretNumber = rand.Intn(10)
	fmt.Println("Secret Number is:", secretNumber)
	fmt.Printf("please enter a number: ")
	fmt.Scan(&userInput)
	fmt.Println("You entered:", userInput)
}

