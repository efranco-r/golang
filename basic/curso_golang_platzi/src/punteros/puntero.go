package puntero

import "fmt"

type PC struct {
	Ram   int
	Disk  int
	Brand string
}

//Personalización de la salida de información
//Sobreescritura del método String
func (myPC PC) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s\n", myPC.Ram, myPC.Disk, myPC.Brand)
}

func (myPC PC) Ping() {
	fmt.Println(myPC.Brand, "Pong")
}

func (myPC *PC) DuplicateRAM() {
	myPC.Ram *= 2
}
