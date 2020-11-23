package main

import "fmt"

func main(){
	// El valor de los slices por defecto es -> null
	// Declaracion de slices
	array := [] int {1,2,3,4,66}

	// Devuelve los dos primeros indices
	/*
	Nos devuelve un array con los valores que le indicamos al slice
	*/
	slice := array[:2]

	// Pintamos el slice
	fmt.Println(slice)
}