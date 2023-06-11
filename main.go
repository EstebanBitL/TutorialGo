package main

import (
	"fmt"
	"tutorialGo/variables"
	//go/src/tutorialGo/variables
)

func main() {
	//variables.MuestroVariables()
	//variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(1566)
	fmt.Println(estado, texto)
}
