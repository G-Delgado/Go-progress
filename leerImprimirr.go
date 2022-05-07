package main

import (
	"bufio"
	"fmt"
	"os"
)

func main4() {
	edad := 20
	// Hay un salto de línea ahí
	// %d le indica a fmt que el valor es un int
	fmt.Printf("Mi edad es %v\n", edad)

	// %t creo que es para booleanos

	// Para flotantes tenemos: %f (Numeros más pequeños)
	// %e flotantes más cientificos
	// %s flotantes más cientificos

	// Si usamos %v, significa que será un valor por defecto

	// Averiguar sobre formatos %v,   %w,   %s,  %t,   %d

	var edad2 int
	fmt.Println("Coloca tu edad: ")
	// Agregamos un salto de línea para que no lea el enter que le damos
	fmt.Scanf("%d\n", &edad2)
	// & El ampersand es un puntero para la variable a la que se lo asigna
	fmt.Printf("Mi edad es: %d\n", edad2)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	// Se coloca el salto de línea pues ahí termina la lectura
	// Cuando lea el salto de línea
	text, err := reader.ReadString('\n')
	// El error debe de ser igual a nil. Puesto que
	// Esta función nos retorna si hay algún error
	// Por tanto si error es nil, es que no hay errores
	// Por tanto debe ser != nil
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hola " + text)
		fmt.Printf("Hola %s", text)
	}
}
