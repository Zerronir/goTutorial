package main

import "fmt"

func main(){
	x := 5
	y := 6

	if(x > y) {
		fmt.Printf("%d es mayor que %d", x, y)
	} else if(x < y){
		fmt.Printf("%d es menor que %d", x, y)
	}

}
