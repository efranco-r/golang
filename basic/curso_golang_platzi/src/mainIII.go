package main

import "fmt"

func mainIII() {
	//Cuando declaramos los tipos de variables, ganamos performance
	//Para los int, ver la clase builtin.go ubicada en \Go\src\builtin
	//Para guardar valores positivos usar uintXX donde XX puede tomar los valores 8, 16, 32 y 64

	//Usando FMT Package
	helloMessage := "Hello"
	worldMessage := "world"

	//Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "Platzi"
	var cursos uint16 = 500

	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	//Usar %v cuando no sabemos el tipo de dato a imprimir
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	//Sprintf, = No imprime en consulta, lo que hace es que guarda la cadena en un String
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Conocer el Tipo de datos de una variable
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T", cursos)
}
