package main


import (
	"fmt"
	"syscall"
)

func main() {

	cwd, _:= syscall.Getwd()

	fmt.Println("Current Working Directory: ", cwd)
}
