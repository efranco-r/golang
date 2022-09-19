package main

import (
	pk "curso_golang_platzi/src/mypackage"
	"fmt"
)

func mainXIII() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	/*
		var myOtherCar pk.carPrivate
		fmt.Println(myOtherCar)
	*/

	//Método público
	pk.PrintMessage("Hola Platzi")
	//Método privado
	//pk.printMessageX("Hola Platzi")
}
