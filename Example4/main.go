package main

import "fmt"

func add(x int, y int) int {
	var sum=x + y
	return sum
}

func main() {
	fmt.Println(add(42, 13))
}
