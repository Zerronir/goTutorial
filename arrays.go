package main

import "fmt"

func main() {
	//var array1 [10] int

	// Definimos una array con longitud y valores
	array2 := [3] int {0, 100, 23}
	// Recorremos el array en go
	for i:=0; i < len(array2); i++ {
		fmt.Println(array2[i])
	}

}