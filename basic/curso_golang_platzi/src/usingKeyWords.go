package main

import "fmt"

func mainVIII() {
	//DEFER: Antes de morir la función, se ejecutar la línea marcada, esto puede verse reflejado como un cambio en el orden de ejecución
	//El comando DEFER es muy usado para cerrar conexiones a BD o cerrar Buffers de Información (FILES), entre otros.
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//CONTINUE y BREAK: Estos son muy usados en los ciclos
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 2 {
			fmt.Println("Es 2")
			continue
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 2 {
			fmt.Println("Es 2")
			break
		}
	}

	//Probando el LABELING en Go
	fmt.Println("---------------------------------------------")

	entry1 := []string{"Jack", "John", "Jones"}
	entry2 := []string{"Peter", "Jack", "Elizabeth"}
	/*
	   MYLABEL:
	   	for _, name1 := range entry1 {
	   		fmt.Printf("Name1: %s\n", name1)

	   		for _, name2 := range entry2 {
	   			fmt.Printf("Name2: %s\n", name2)
	   			if name2 == name1 {
	   				fmt.Printf("Name2: %s is equals Name1: %s\n", name2, name1)
	   				continue MYLABEL
	   			}

	   			fmt.Println("Continue for II........")
	   		}

	   		fmt.Println("Continue for I........")
	   	}
	*/
MYLABEL:
	for _, name1 := range entry1 {
		fmt.Printf("Name1: %s\n", name1)

		for _, name2 := range entry2 {
			fmt.Printf("Name2: %s\n", name2)
			if name2 == name1 {
				fmt.Printf("Name2: %s is equals Name1: %s\n", name2, name1)
				break MYLABEL
			}

			fmt.Println("Continue for II........")
		}

		fmt.Println("Continue for I........")
	}
}
