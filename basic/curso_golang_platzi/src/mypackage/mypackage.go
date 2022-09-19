package mypackage

import "fmt"

//CarPublic Car con acceso publico
type CarPublic struct {
	//Cuando utilizamos la primera letra en MAYUSCULAS
	//Estamos diciendo que la estrucutura o variables son PUBLICAS
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

//PrinteMessage Imprimer un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}

//PrinteMessage Imprimer un mensaje (PRIVADO)
func printMessageX(text string) {
	fmt.Println(text)
}
