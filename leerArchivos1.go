package main

import (
	"fmt"
	"io/ioutil"
)

func main17() {
	// El problema de esta forma de leer es que hay que esperar que
	// Se lea todo el archivo primero
	/*
		Por lo que si es muy pesado, puede tardar mucho tiempo
	*/

	// El lee el archivo desde donde se genera el binario
	// Así que si corremos go desde la carpeta por fuera, da error
	data, err := ioutil.ReadFile("./hola.txt")

	if err != nil {
		fmt.Println(err)
	}
	// Se encuentra en bytes
	fmt.Println(data)
	// AHORA en string
	fmt.Println(string(data))
}
