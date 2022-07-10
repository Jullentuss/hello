package main

import (
	"fmt"

	"./greet"
)

func main() {
	// en los paquetes, para que las variables o funciones sean publicas
	// su nombre debe inicar con mayuscula
	// de lo contrario son privadas
	// las variables declaradas a nivel de paquetes, son accesibles desde todos los paquetes del
	v := greet.Spanish()
	fmt.Println(v)
}
