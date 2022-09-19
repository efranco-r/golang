package main

import "fmt"

//Una buena practica al momento de declarar un parametro al  método y este es del tipo CHANNEL, lo ideal es
//Determinar si ese parámetro es de entrada o de salida
//chan<- Determina que el Channel es de entrada
//<-chan Determina que el Channel es de salida
func saySomething(text string, c chan<- string) {
	//Tomaremos el texto que se recibe y se lo pasaremos al channel
	c <- text
	//fmt.Println(texto)
}

func mainXIX() {
	//Los CHANNELS permiten compartir datos entre los Goroutines, ya que ellos manejan de forma nativa entre ellos y otros datos primitivos
	//Un CHANNEL = KERNEL
	//Determino el tipo de dato que pasare por el Channel y ademas la cantidad de datos simultaneos que manejará el canal
	//Es una buena practica indicarle la cantidad de datos simultaneos que manejará, en este caso, 1
	c := make(chan string, 1)

	fmt.Println("Hello")

	//Invocamos el método con una Goroutine
	go saySomething("Bye", c)

	//Para hacer que el hilo principal MAIN espere por la ejecución de la Goroutine en el Channel, hacemos lo siguiente
	fmt.Println(<-c)
}
