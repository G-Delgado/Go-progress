package main

import "fmt"

// No es necesario importar nada pues comparten el paquete, toncs User struct ya está
// Retorna un string
// GO no es orientado a objetos y por eso no tiene clases
// Pero esto tiene algunos tintes
func ( /*usuario*/ this User) nombre_completo() string {
	// return usuario.nombre + " " + usuario.apellido
	return this.nombre + " " + this.apellido
}

/*
	Cuando recibimos la estructura como un puntero en lugar de una copia
	Nos permite cambiar el objeto original
*/

// Intento de setter
// Esto no cambia el nombre, no funciona
// La estructura que está modificando es una copia
// No es el original.
func (this User) set_name(n string) {
	this.nombre = n
}

// Pasar una referencia es más eficiente que pasar una copia
func (this *User) set_name_real(n string) {
	this.nombre = n
}

// Si importamos estructuras de otros paquetes no podemos añadirle
// Las funcionalidades o métodos de arriba

func main13() {
	gabriel := new(User)

	gabriel.nombre = "Gabriel"
	gabriel.apellido = "Alejandro"

	// No funciona
	gabriel.set_name("Brand")
	// Funciona
	gabriel.set_name_real("Brand")

	fmt.Println(gabriel.nombre_completo())
}
