package main

import "fmt"

func main() {
	// Podemos declarar arrays de esta manera y hacerlas dinámicas
	slice := make([]int, 3, 5)

	// podemos usar cap() para la capacidad de la array

	fmt.Println("Capacidad -> ", cap(slice))
	fmt.Println("Longitud -> ", len(slice))
	fmt.Println("Slice -> ", slice)
}