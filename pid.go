package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("what up")
}
func Getpid(urxvt) int {
	return syscall.Getpid(urxvt)
	fmt.Println(urxvt)

}
