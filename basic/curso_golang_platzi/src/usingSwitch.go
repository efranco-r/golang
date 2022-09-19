package main

import "fmt"

func mainVII() {
	//Switch con condiciones
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//Switch sin condiciones
	//Se usa cuando se desea anidar multiples condiciones como cuando usas multiples IFs
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condición")
	}
}
