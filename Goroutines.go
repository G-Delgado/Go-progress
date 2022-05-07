package main

import (
	"fmt"
	"strings"
	"time"
)

func main16() {
	// go mi_nombre_lento("Gabriel")
	// // La linea de abajo tiene que esperar a que termine la de arriba
	// fmt.Println("Que aburrido esperar!")
	// Eso se puede arreglar poniendole la palabra clave
	// go mi_nombre_lento("Gabriel")
	// De esa forma se llama concurrentemente

	// Sin embargo, el programa terminó sin imprimir el nombre

	go mi_nombre_lento("Gabriel")

	fmt.Println("Que aburrido esperar!")

	// Así el programa no se acaba hasta que le de un enter
	var wait string
	fmt.Scanln(&wait)

	/*
		Puedes ponerlo en funciones anonimas
		go func() {

		}()
	*/
}

func mi_nombre_lento(name string) {
	// "hola_mundo".Split("hola_mundo", "_") = hola mundo
	letras := strings.Split(name, "")

	// El primer parametro es el index, pero no nos importa
	for _, letra := range letras {
		fmt.Println(letra)
		time.Sleep(1000 * time.Millisecond)
	}
}
