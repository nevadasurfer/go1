package main

import (
	"fmt"
	"math"
)

func f1() {
	fmt.Println("this is not main")
}

func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

func f4(a float64) float64 {
	return math.Pow(a, a)
}

func f5(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	f3(345, 345, 3, 4.3, 5.6, "whatUP")

	p := f4(5.3)
	fmt.Println(p)

	a, b := f5(8, 9)
	fmt.Println(a, b)
}
