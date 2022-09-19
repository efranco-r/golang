package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	//Esto lo hacemos para librerar la goroutine una vez el método say termine
	defer wg.Done()

	fmt.Println(text)
}

func mainXVIII() {
	//Creamos una Whilegroup
	//Los whilegroups son mas eficientes que manejar CHANNELS
	var wg sync.WaitGroup

	fmt.Println("Hello")

	//Esto nos indica que vamos a agregar una goroutine al whilegroup
	wg.Add(1)
	//Para ejecutar en forma concurrente solo usamos la palabra reservada "go"
	go say("world", &wg)
	//Esto lo usamos para decirle a la gourutine principal (main) que espere
	wg.Wait()

	//Esto no es eficienta, es usado solo para demostrar la ejecución concurrente del segundo Print
	//time.Sleep(time.Second * 1)

	/*
	* OJO: Las Goroutines son muy usadas con funciones anonimas
	 */

	//Ejemplo de como aplicar una función anonima con goroutines
	go func(text string) {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second * 1)
}
