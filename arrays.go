package main

import "fmt"

func main() {
	// Por defecto cogen un valor 0, vacío o nulo
	//var array1 [10] int

	// Definimos una array con longitud y valores
	array2 := [3] int {0, 100, 23}
	// Recorremos el array en go
	for i:=0; i < len(array2); i++ {
		fmt.Println(array2[i])
	}

	// Arrays bidimensionales
	var matrix[2][2] int
	fmt.Println(matrix)
}