package main

import "fmt"

func main() {
	// Solo existe el ciclo for en go

	//1. Bucle básico
	/*for i:=0; i<10; i++ {
		fmt.Printf("Hello World \n")
	}*/

	for i := 0; i < 10; i += 2 {
		fmt.Println(i)
	}

	// Tabla de multiplicar
	for i := 0; i < 10; i++ {
		fmt.Println(i * 3)
	}

	// Se puede simular un while de la siguiente manera:
	y := 0
	for y < 10 {
		fmt.Println(y)
		y++
	}

	// Para salir de un ciclo podemos usar --break--

}