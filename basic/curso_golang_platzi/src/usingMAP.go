package main

import "fmt"

func mainXI() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	//Recorrer un MAP
	/* Cuando usar MAPs
	 * 1. Cuando tengo 2 o mas listas que se relacionan entre si
	 * 2. Los MAP son mas eficientes que manejar SLICE o ARRAYs ya que de forma nativa implementan CONCURRENCIA
	 */
	for index, value := range m {
		fmt.Println(index, value)
	}

	//Obtener un valos por su KEY
	value := m["Jose"]
	fmt.Println("Key:", value)

	//¿Que pasa si colocamos una llave que no existe?
	value = m["Josep"]
	fmt.Println("Key:", value) //Muestra el Zer-Value del tipo

	//¿Como hacemos para tener certeza al momento de buscar una llave que posiblemente no este en el DICCIONARIO?
	//La variable OK nos pemite mostrar si la llave que estamos buscando existe o no en el DOCCIONARIO
	key, ok := m["Josep"]
	fmt.Printf("Key: %d Estatus: %v\n", key, ok) //Muestra el Zer-Value del tipo

}
