package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Sqrt(25))
	fmt.Println(math.Sqrt(225))
	fmt.Println(math.Sqrt(64.25))
	fmt.Println(math.Sqrt(123.45))

	fmt.Println(math.Sqrt(1))
	fmt.Println(math.Sqrt(0))

	fmt.Println(math.Sqrt(-1))
	// generating infinite value
	fmt.Println(math.Sqrt(math.Inf(-1)))
	fmt.Println(math.Sqrt(math.Inf(1)))
	fmt.Println(math.Sqrt(math.NaN()))
}