package main

import "fmt"

func main2() {
	// Se pueden generar varias variables y asignares el tipo directamente
	// Se inician en 0
	var x, y, z int
	// Se inician en ""
	var cadena string
	// Se inician en false
	var bandera bool

	// Go es muy estricto. No podemos tener variables declaradas sin
	// usar.

	// Este operador declara e inicializa el valor directamente
	g := 23

	// --------------------------

	nombre := "Coco"
	// No podemos hacer esto, ya que estamos volviendo a declarar la misma
	// nombre := "Cocooo"

	// Para que no me den error ;'v
	fmt.Println(nombre, g, x, y, z, cadena, bandera)

}
