package main


import (
	"fmt"
	"syscall"
)

func main() {

	err := syscall.Mkdir("/home/jkennedy/devname",0754)

	if err == nil {

		fmt.Println("Directory Created")
	}
}
