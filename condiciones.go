package main

import "fmt"

func main5() {
	// if
	// else if
	// else

	// En go no es necesario usar paréntesis para los condicionales
	// Valores booleanos : true, false5
	// Las llaves de los condicionales no pueden ir abajo
	// Si no en la misma línea
	// Tampoco podemos poner el if sin llaves así sea solo una línea

	x := 10
	y := 5
	if x > y {
		fmt.Printf("%d es mayor que %d\n", x, y)
	} else if x < y {
		fmt.Printf("%d es menor que %d\n", x, y)
	} else {
		fmt.Printf("%d es igual que %d\n", x, y)
	}

	// Podemos colocar n else if

	/*
		Operadores lógicos:
		==	Igual a
		!=	Diferente de
		<	Menor que
		>	Mayor que
		>=	Mayor o igual que
		<=	Menor o igual que
		&&	AND
		||	OR
	*/

}
