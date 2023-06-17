package main

import (
	//"tutorialGo/variables"

	"tutorialGo/files"
)

func main() {

	//variables.MuestroVariables()
	//variables.RestoVariables()
	//estado, texto := variables.ConviertoaTexto(1566)
	//fmt.Println(estado, texto)
	/*if os := runtime.GOOS; os == "linux" || os == "windows" {
		fmt.Println("Esta en linux")
	} else {
		fmt.Println("Esta en otro sistema operativo")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esta en linux")
	case "darwin":
		fmt.Println("Esta en darwin")
	default:
		fmt.Println("%s \n", os)
	}

	//numero, texto := ejercicios.ConvNumericos("holaa")
	//fmt.Println(numero, texto)

	teclado.IngresarNumero()

	iteraciones.Iterar()

	fmt.Println(ejercicios.TabladeMultiplicar())*/

	//files.GrabaTabla()

	//files.SumaTabla()

	files.LeerArchivo()
}
