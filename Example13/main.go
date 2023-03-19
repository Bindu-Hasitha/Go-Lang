package main

import "fmt"

func main() {
	var (
		firstName, lastName string
		age                 int
		salary              float64
		isConfirmed         bool
	)

	fmt.Printf("firstName: %s, lastName: %s, age: %d, salary: %f, isConfirmed: %t\n",
		firstName, lastName, age, salary, isConfirmed)
	fmt.Printf("types of variables are: %T %T %T %T %T",firstName, lastName,age,salary,isConfirmed)
}