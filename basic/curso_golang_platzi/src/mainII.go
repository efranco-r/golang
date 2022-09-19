package main

import (
	"fmt"
	"math"
)

func mainII() {
	//Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	//Suma
	resultado := x + y
	fmt.Println("Suma: ", resultado)

	//Resta
	resultado = y - x
	fmt.Println("Resta: ", resultado)

	//Multiplicación
	resultado = x * y
	fmt.Println("Multiplicación: ", resultado)

	//División
	resultado = y / x
	fmt.Println("División: ", resultado)

	//Modulo
	resultado = y % x
	fmt.Println("Modulo: ", resultado)

	//Incremental
	x++
	fmt.Println("Incremental: ", x)

	//Decremental
	x--
	fmt.Println("Decremento: ", x)

	//Reto
	//Calcular el area de un rectangulo, trapecio y de un circulo

	//Area de un rectangulo
	var baseRectangulo int = 10
	var alturaRectangulo int = 5

	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Printf("El area de un rectangulo con base %d y altura %d es: %d", baseRectangulo, alturaRectangulo, areaRectangulo)

	//Area de un trapecio
	var baseA int = 6
	var baseB int = 3
	var alturaTrapecio int = 4

	areaTrapecio := alturaTrapecio * (baseA + baseB) / 2
	fmt.Printf("\nEl area de un trapecio con altura %d y base A %d con base B %d es: %d", alturaTrapecio, baseA, baseB, areaTrapecio)

	//Area de in circulo
	const pi float64 = math.Pi
	var radio float64 = 0.5

	areaCirculo := pi * math.Pow(radio, 2)
	fmt.Printf("\nEl area de un circulo con radio %g es: %g", radio, math.Trunc(areaCirculo))
	fmt.Printf("\nEl area de un circulo con radio %g es: %g", radio, math.Abs(areaCirculo))
	fmt.Printf("\nEl area de un circulo con radio %g es: %g", radio, math.Round(areaCirculo))
	//Esta salida me permite redondear el valor a su proximo valor flotante
	fmt.Printf("\nEl area de un circulo con radio %g es: %.2f", radio, areaCirculo)

}
