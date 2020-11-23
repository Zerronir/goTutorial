package main

import (
	"fmt"
	"strconv"
)

func main(){
	// Convertir String a int
	edad := "22"
	edad_int,_ := strconv.Atoi(edad)
	fmt.Println(edad_int + 10)
}