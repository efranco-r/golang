package main

import "fmt"

type shape2D interface {
	area() float64
}

type square struct {
	base float64
}

type rectangle struct {
	base   float64
	height float64
}

func (mySquare square) area() float64 {
	return mySquare.base * mySquare.base
}

func (myRectangle rectangle) area() float64 {
	return myRectangle.base * myRectangle.height
}

func calculate(myShape2D shape2D) {
	fmt.Println("Area:", myShape2D.area())
}

func mainXVII() {
	mySquare := square{base: 2}
	myRectangle := rectangle{base: 2, height: 4}

	//Without interfaces
	fmt.Println(mySquare.area())
	fmt.Println(myRectangle.area())

	//With interfaces
	calculate(mySquare)
	calculate(myRectangle)

	//Using interfaces list
	//[1, "bo", true, ]//This is not posible in GOLANG
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)
}
