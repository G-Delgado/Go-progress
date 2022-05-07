package main

import "fmt"

func main8() {
	// Los slices son más comunes porque nos permite redimensionar
	// Los 'arreglos' en tiempo de ejecución
	// No tenemos que definir el tamaño en tiempo de ejecución

	// [3]int es diferente a [5]int
	// Se declara un slice...
	var slice []int
	fmt.Println(slice)
	// Se inicializan en nulo / nil / null
	// En GO no es muy común el nil, puesto que las variables
	// Tienen un valor "0" al declararse
	fmt.Println(slice == nil)

	slice2 := []int{1, 2, 3, 4}
	fmt.Println(slice2)

	// len(slice) también es posible

	// Puntero al arreglo
	// Longitud del arreglo al que apunta
	// Capacidad

	// Obtener slice a partir de un arreglo
	arreglo := [3]int{1, 2, 3}
	slice3 := arreglo[:2]

	// También se puede hacer otros slices [1:3] Devuelve 2 y 3
	// [:3] Toma el valor inicial, osea desde el inicio hasta el valor final asignado
	// [1:] Nos devuelve de esa posición al final
	fmt.Println(slice3)
}
