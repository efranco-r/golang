package main

import "fmt"

func mainI() {
	//Declaración de CONSTANTES & VARIABLES
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Printf("Pi1: %g\nPi2: %g", pi, pi2)

	//Declaración de variables enteras
	//:= Define que la variable no ha sido declara con anterioridad

	//Utilizar la forma corta para declaración e inicio de la variable
	base := 12
	//Utilizar la forma larga, cuando no inicialice la variable
	var altura int = 4
	var area int

	fmt.Printf("\nBase: %d\nAltura: %d\nArea: %d", base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Printf("\nInt: %d\nFloat: %g\nString: %s\nBoolean: %t", a, b, c, d)

	//Calcular el area de un cuadrado
	const baseCuadrado = 5
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Printf("\nAltura: %d\nBase: %d\nArea: %d", baseCuadrado, baseCuadrado, areaCuadrado)

	//Revisar esto, como manejar las diferentes formas de imprimir:
	//https://www.geeksforgeeks.org/difference-printf-sprintf-fprintf/
}
