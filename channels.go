package main

import "fmt"

func main18() {
	// Los channels nos permiten comunicar Goroutines unas con otras
	channel := make(chan string)

	go func(channel chan string) {
		for {
			var name string
			fmt.Scanln(&name)
			// Enviar info a un canal
			// Canal <- Información a enviar
			channel <- name
		}
	}(channel)

	// Recibir información
	// var msg string
	// msg = <- channel
	msg := <-channel

	fmt.Println(msg)

	// De esta forma, luego de recibir información tiene que volver a pedirla
	// Por eso nos entra en el bucle infinito.

	msg = <-channel

	fmt.Println("Esta es la segunda vez que recibo información: " + msg)
}
