package main

import "fmt"

func main11() {
	/*
		1. Un puntero es una dirección de memorria
		2.En lugar del valor, tenemos la dirección en la que está el valor.
		3. X, Y => #as123d => 5
		Digamos que X y Y apuntan a esa misma dirección de memoria
		4. X => 6. Entonces...
		5. X => #as123d => 6
		6. Y ¿? => 6
		Ya ahora va a ser 6 porque X modificó el valor al que Y está apuntando.
		*T => *int, *string, *Struct
		Declaramos asterisco y el tipo de variable
		Valor zero => nil
	*/

	var x, y *int
	entero := 5
	// En lugar de x = 5
	// Así accede a la dirección del dato de la variable entero
	x = &entero
	y = &entero

	fmt.Println(x)
	fmt.Println(y)

	// Se pone asterico x porque x es la dirección
	// Entonces el puntero apunta al valor en esa dirección
	*x = 6

	fmt.Println(*x)
	fmt.Println(*y)

}
