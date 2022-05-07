package main

import "fmt"

func main7() {

	// var arreglo [/*Cantidad de elementos que caben en el arreglo*/]
	// No podemos meter diferentes tipos en un mismo arreglo
	var arreglo [5]int
	fmt.Println(arreglo)
	arreglo2 := [3]int{1, 2, 3}
	fmt.Println(arreglo2)
	arreglo3 := [3]int{1, 2}
	fmt.Println(arreglo3)

	fmt.Println(len(arreglo))
	// Sirve el len arreglo para iteraciones
	// También podemos acceder a sus indices arreglo[i]

	// También podemos modificar arreglo[2] = 20

	// MATRICES
	// El primer corchete es filas y el segundo columnas
	// Se iteran de la misma forma que en otros lenguajes [i][j]
	var matriz [3][2]int
	fmt.Println(matriz)

}
