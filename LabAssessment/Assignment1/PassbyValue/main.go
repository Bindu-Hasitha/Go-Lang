package main

import "fmt"

type Student struct {
    Name string
    Branch  string
	Marks int
}

func changeMarks(s Student) {
    s.Marks = 80
}

func main() {
    student := Student{Name: "Bindu", Marks: 90}
    fmt.Println(student) // Output: {Bindu 90}

    changeMarks(student)
    fmt.Println(student) // Output: {Bindu 90}
}

//  it is passed by value, which means that a copy of the struct is made and passed to the function. 
//Any changes made to the struct inside the function do not affect the original struct.