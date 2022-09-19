package main

import (
	"fmt"
	"strings"
)

// RETO Construir una función que reciba un parametro y que identifique si el número es PAR o IMPAR
func numberEvenOrOdd(value int) string {
	var result string = "Impar"

	//if number%2 == 0 {
	if module := value % 2; module == 0 {
		result = "Par"
	}

	return result
}

// RETO Una función que reciba un usuario y una contraseña y que valide si el usuario es valido o no
func isCredentialAccept(user, password string) bool {
	isAccept := false
	if len(strings.TrimSpace(user)) != 0 && len(strings.TrimSpace(password)) != 0 && user != password {
		isAccept = true
	}

	return isAccept
}

func mainV() {
	/*
		valor1 := 2

		if valor1 == 1 {
			fmt.Println("Valor es:", valor1)
		} else {
			fmt.Println("Es otro valor:", valor1)
		}

		//Convertir texto a número
		//value, err := strconv.Atoi("53")
		value, err := strconv.Atoi("strfffrr")

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Value:", value)
	*/
	number := 0
	result := numberEvenOrOdd(number)
	fmt.Printf("El número %d es: %s\n", number, result)

	user := "eduardo"
	password := "franco"
	isAccept := isCredentialAccept(user, password)
	fmt.Printf("El usuario '%s' con contrasseña '%s' es aceptado: %v", user, password, isAccept)

}
