package main

import (
	"fmt"
	"math/cmplx"
)

var (
	boolean   bool       = false
	number int    = 20
	unit uint64= 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", boolean, boolean)
	fmt.Printf("Type: %T Value: %v\n", number, number)
	fmt.Printf("Type: %T value: %v\n",unit,unit)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}