package main

import "fmt"

func normalFunction(message string) {
	fmt.Printf("My function %s\n", message)
}

func tripeArgumentsI(a int, b int, c string) {
	fmt.Printf("a %d b %d c %s\n", a, b, c)
}

//Esto es una buena practica en GO
//GO no permite sobre carga de constructores
func tripeArgumentsII(a, b, c int, d string) {
	fmt.Printf("a %d b %d c %d d %s\n", a, b, c, d)
}

func returnValue(a int8) int8 {
	return a * 2
}

func returnDoubleValue(a int8) (b, c int8) {
	return a, a * 2
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func mainIV() {
	normalFunction("1")
	tripeArgumentsI(2, 4, "Eduardo")
	tripeArgumentsII(2, 4, 5, "Eduardo")

	var value int8 = returnValue(2)
	fmt.Println("Value:", value)

	var value1, value2 int8 = returnDoubleValue(2)
	fmt.Printf("Value1: %d Value2: %d\n", value1, value2)

	//Si en un momento un método me regresa mas de dos valores, pero solo me interesa uno
	value3, _ := returnDoubleValue(3)
	fmt.Println("Value3", value3)

	fmt.Println(split(2))
}
