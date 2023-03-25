package main

import "fmt"

type Student struct {
	Name   string
	Branch string
	Marks  int
}

func changeAge(s *Student) {
	s.Marks = 80
}

func main() {
	student := &Student{Name: "Bindu", Marks: 90}
	fmt.Println(*student) // Output: {Bindu 90}

	changeAge(student)
	fmt.Println(*student) // Output: {Bindu 80}
}

//When you pass a pointer to a struct, the function receives a copy of the pointer, which points to the same memory location as the original struct.
//Any changes made to the struct inside the function will affect the original struct.
