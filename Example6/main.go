package main
 
import (
	"fmt"
)

func calculator(a,b int) (mul int,div int){
	mul=a*b
	div= a/b
	return
}

func main(){
	m,d :=calculator(100,50)
	fmt.Println("100*50=",m)
	fmt.Println("100/50=",d)
}