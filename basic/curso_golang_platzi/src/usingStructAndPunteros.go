package main

import (
	puntero "curso_golang_platzi/src/punteros"
	"fmt"
)

func mainXIV() {
	a := 50
	//b va a ser la dirección de memoria donde se esta guardando a
	b := &a

	fmt.Println(b)
	//Se usa * para acceder al valor de esa dirección en memoria
	fmt.Println(*b)

	a = 100
	fmt.Println(a, *b)

	*b = 200
	fmt.Println(a, *b)

	//¿En que se usan los punteros en go?
	//No solo para construir funciones un poco mas costumisadas
	//Si no tambien para poder llevar las funciones de una libreria a otro paquete de forma mucho mas facil y eficiente

	//myPc := punteros.PC{ram: 16, disk: 200, brand: "msi"}
	var myPc puntero.PC = puntero.PC{Ram: 16, Disk: 200, Brand: "msi"}
	fmt.Println(myPc)

	myPc.Ping()

	//USO DE PUNTEROS, para acceder a memoria y modificar carcateristicas del STRUCT
	//Primera forma
	myPc.Ram = 20

	//Segunda forma usando punteros
	myPc.DuplicateRAM()
	fmt.Println(myPc)
}
