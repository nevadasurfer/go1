package main

import (
	"fmt"
	"syscall"
)

func main() {

	env := syscall.Environ()

	for i := range env {
		fmt.Println(env[i])
	}
}
