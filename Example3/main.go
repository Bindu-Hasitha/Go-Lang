package main

import (
	"fmt"
	"math"
)

func area(r float64) float64{
	var pi float64=math.Pi
	var ar float64 =2*pi*r
	return ar
}

func main(){
	
	
	fmt.Println("Area of circle is",area(10))
}