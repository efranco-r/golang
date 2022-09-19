package main

import (
	"fmt"
	"strings"
)

func esPalindromo(word string) bool {
	esPalindromo := true
	sliceWord := []rune(strings.ToLower(word))

	indiceSliceWord := len(sliceWord) - 1

	for _, valor := range sliceWord {
		charIzquierdaDerecha := string(valor)

		if charDerechaIzquierda := string(sliceWord[indiceSliceWord]); charIzquierdaDerecha == charDerechaIzquierda {
			indiceSliceWord--
		} else {
			esPalindromo = false
			break
		}
	}

	return esPalindromo
}

func esPalindromoPlatzi(text string) {
	var textReverse string

	for index := len(text) - 1; index >= 0; index-- {
		textReverse += string(text[index])
	}

	if strings.ToLower(text) == strings.ToLower(textReverse) {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func mainX() {
	slice := []string{"hola", "que", "pasa"}

	for indice, valor := range slice {
		fmt.Println(indice, valor)
	}

	//Si no me interesa tener el indice, lo escapamos con _
	for _, valor := range slice {
		fmt.Println(valor)
	}

	//Si solo me interesa saber el indice
	for indice := range slice {
		fmt.Println(indice)
	}

	//RETO: Determinar si una palabra es un PALINDROMO
	//Ejemplo 1: ama
	//Ejemplo 2: amor a roma
	word := "Amor a roma"

	//startP1 := time.Now()
	//fmt.Println(startP1)
	fmt.Printf("La oración '%s' es un palindromo %v\n", word, esPalindromo(word))
	//fmt.Println("Tiempo estimado esPalindromoI:", time.Since(startP1))

	//fmt.Printf("Tiempo estimado esPalindromoI: %v\n", endP1.Nanosecond()-startP1.Nanosecond())

	//startP2 := time.Now()
	//fmt.Println(startP2)
	esPalindromoPlatzi(word)
	//fmt.Println("Tiempo estimado esPalindromoII:", time.Since(startP2))
	//Medir el tiempo de ejecución del cuerpo (FIN)
	//endP2 := time.Now()
	//fmt.Printf("Tiempo estimado esPalindromoII: %v\n", endP2.Nanosecond()-startP2.Nanosecond())
}
