package main

import "fmt"

//Crear un tipo Struct
type car struct {
	brand string
	year  int
}

func mainXII() {
	//Instancias un STRUCT: Forma 1
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	//Instancias un STRUCT: Forma 2
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
