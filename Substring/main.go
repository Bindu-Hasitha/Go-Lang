package main

import (
	"fmt"
	"strings"
)



func main(){
	var a = [5] string {"bin","bindu","duHabinsitha","duhasivbinasireddy","Vijaya"}

	for i:=1; i<len(a);i++{
		if strings.Index(a[i],a[0]) >= 0 {
            fmt.Printf("%s is a substring of %s\n", a[0], a[i])
		}else{
			fmt.Printf("%s is not a substring of %s\n",a[0], a[i])
		}
	}
} 