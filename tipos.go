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

	precio := 14.3
	fmt.Printf("El precio es %f", precio)

	nombre := "Raul"
	fmt.Printf("Mi nombre es %s", nombre)
}