package main

import (
	"fmt"
	"os"
)

func main() {
	uid := os.Getuid("root")
	fmt.Println("uid:", uid)
}
