package main

import "fmt"

func main9() {
	slice := make([]int, 3, 5)
	// make(tipo de slice, longitud, capacidad)
	// Se puede string, boolean, etc
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	// Cap es el capacity del slice.
	// Si no ponemos el tercer arguumento de make, el capacity y longitud son iguales

	// Un slice apunta a un arreglo interno
	// El tamaño de ese arreglo define el len del slice
	// La capacidad es cuantos elementos caben en ese slice

	slice = append(slice, 2)
	fmt.Println(slice)
	fmt.Println(len(slice))

	slice2 := make([]int, 0)
	fmt.Println(slice2)

	// Cuando un slice desborda su capacidad
	// Se construye un nuevo slice con una capacidad más grande
	// Por tanto es más eficiente tener capacidad para no reconstruir slices

	slice3 := make([]int, 0, 8)
	fmt.Println(slice3)

}
