package main

import "fmt"

func mainVI() {
	//GO solo maneja una estructura para realizar ciclos, y es FOR

	//For condicional
	for indice := 0; indice < 5; indice++ {
		fmt.Println("Indice:", indice)
	}

	//For while: se usa para realizar cierta cantidad de pasos hasta que una condición se cumpla
	counter := 0
	for counter < 5 {
		fmt.Println("Counter:", counter)
		counter++
	}

	//For for ever: Indica que el clico iterar hasta la eternidad
	/*
		counterForever := 0
		for {
			fmt.Println("counterForever:", counterForever)
			counterForever++
		}
	*/

	//For  range
	entry := []string{"Jack", "John", "Jones"}
	for i, val := range entry {
		fmt.Printf("At position %d, the character %s is present\n", i, val)
	}
}
