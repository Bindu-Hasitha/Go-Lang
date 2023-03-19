package main
import "fmt"

const Pi = 3.14

func main() {
	const name = "John"
	fmt.Println("Hello", name)
	fmt.Println("Value of Pi is", Pi)

	const myval = true
	fmt.Println("Variable myval contains:", myval)
}