package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	/*var nombre string
	fmt.Println("Dime tu nombre")
	fmt.Scanf("%s \n", &nombre)
	var edad int
	fmt.Println("Cuantos años tienes?")
	// Agregar un & para asignar el valor a la variable
	fmt.Scanf("%d \n", &edad)
	fmt.Printf("%s tiene %d años", nombre, edad)*/

	// Esto devuelve un error por defecto siguiente linea bloque será el corregido
	/*reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	text,err := reader.ReadString('\n')
	if(err == nil) {
		fmt.Println(err)
	} else {
		fmt.Println("Hola %s", text)
	}*/

	// Esto no da error
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	text,err := reader.ReadString('\n')
	if(err != nil) {
		fmt.Println(err)
	} else {
		fmt.Printf("Hola %s", text)
	}

}