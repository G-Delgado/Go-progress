package main

import "fmt"

func main10() {
	//slice := make([]int, 0, 4)
	slice := []int{1, 2, 3, 4}
	copia := make([]int, 4)

	fmt.Println(slice)
	fmt.Println(copia)
	// copy(destino, fuente)
	// Tengo que tener cuidado con llamar una variable copy y usar funcion copy
	// Porque se sobreescribe la función
	copy(copia, slice)

	fmt.Println(slice)
	fmt.Println(copia)

	copia2 := make([]int, 0)
	/*
	 La función copy copia el mínimo de elementos
	 Es decir, si un arreglo vale 0, va a copiar 0 elementos
	 Pues es el mínimo de ambas longitudes
	 Si le pones un make con 1 de longitud, solo copia un elemento.

	 Una forma de evitarlo es:
	 copia3 := make([]int, len(slice))
	 Y otra es duplicar la capacidad
	 copia4 := make([]int, len(slice), cap(slice)*2)
	*/
	copy(copia2, slice)

	fmt.Println(slice)
	fmt.Println(copia2)
}
