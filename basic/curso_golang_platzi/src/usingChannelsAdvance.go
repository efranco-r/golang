package main

import "fmt"

func message(text string, c chan<- string) {
	c <- text
}

func mainXX() {
	c := make(chan string, 3)
	c <- "Mensaje 1"
	c <- "Mensaje 2"

	//Usamos len() para validar cuantos Goroutines hay dentro del channel
	//Usamos cap() para validar la cantidad máxima de Goroutines que puede almacenar el Channel
	fmt.Printf("Cantidad de GOROUTINE insertadas: %d, Cantidad máxima de GORUOTINES aceptadas: %d\n", len(c), cap(c))

	//Como hacer el recorrido de cada uno de los valores guardados
	//Range y Close

	//Close: Le dice al runtime de GO que va a cerrar el Channel
	//Por buenas practicas, es importante cerrar los Channel una vez que no requieres aplicar mas entradas al mismo
	close(c)
	//La sgte línea genera un error, ya que el Channel ya se encuentra cerrado
	//c <- "Mensaje3"

	//Range: Nos permite recorrer todos los valores que se encuentran en el Channel
	//OJO: Si no he cerrado el channel, me genera una EXCEPCION
	for message := range c {
		fmt.Println(message)
	}

	//Cuando no tenemos certeza de que canal va a responder primero, es cuando debemos usar SELECT
	//Select
	email1 := make(chan string) //Channel 1
	email2 := make(chan string) //Channel 2

	go message("Mensaje para Ingrid", email1)
	go message("Mensaje para Eduardo", email2)

	for index := 0; index < 2; index++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}
}
