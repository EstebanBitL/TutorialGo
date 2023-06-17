package files

import (
	"fmt"
	//"io/ioutil"
	"bufio"
	"os"

	//"text/scanner"
	"tutorialGo/ejercicios"
)

var fileName string = "files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.TabladeMultiplicar()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto) //con Fprintln escribe en el archivo
	archivo.Close()              //una ves que se termina de escribir se cierra el archivo
}

func SumaTabla() {
	var texto string = ejercicios.TabladeMultiplicar()
	if !Append(fileName, texto) {
		fmt.Println("Error al crear el archivo")
	}
}

//con esta funcion se agrega texto al archivo tabla.txt
func Append(fileN, texto string) bool {
	archivo, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error durante el append" + err.Error())
		return false
	}
	_, err = archivo.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el writeString" + err.Error())
		return false
	}
	archivo.Close()
	return true
}

//con esta funcion lee todos los archivos de tabla.txt 
func LeerArchivo() {
	//archivo, err := ioutil.ReadFile(fileName)
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := (scanner.Text())
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
