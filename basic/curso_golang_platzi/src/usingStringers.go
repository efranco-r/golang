package main

import (
	pc "curso_golang_platzi/src/punteros"
	"fmt"
)

// Personalización de salidas de caracteres sin usar Printf
func mainXV() {
	myPC := pc.PC{Ram: 16, Brand: "MSI", Disk: 100}
	fmt.Println(myPC)
}
