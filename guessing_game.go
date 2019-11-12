package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	var userInput int
	var secretNumber int

	rand.Seed(time.Now().Unix())
	secretNumber = rand.Intn(10)
	fmt.Println("Secret Number is:", secretNumber)


	fmt.Printf("please enter a number: ")
	fmt.Scan(&userInput)
	fmt.Println("You entered:", userInput)

	if userInput == secretNumber {
		fmt.Println("You Won!!")
	}	else if userInput < secretNumber {
		fmt.Println("Too Low...")
	}	else if userInput > secretNumber {
		fmt.Println("Too High...")

		}

	}


