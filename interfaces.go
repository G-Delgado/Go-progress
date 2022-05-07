package main

import "fmt"

// Una interfaz en go es una estructura que define metodos vacíos
// También es un tipo de dato que podemos pasar en funciones, arreglos, etc

type Person interface {
	// No estoy seguro de si es en mayusculas o minusculas
	Permissions() int // 1 - 5
	Name() string
}

type Admin struct {
	name string
}

// No hay un keyword implements.
/*
Simplemente si los métodos implementados a Admin tienen los mismos
metodos de la interfaz, está implementada
*/
// ----------------------------------------

func (this Admin) Permissions() int {
	return 5
}

func (this Admin) Name() string {
	return this.name
}

// ----------------------------------------
// A pesar de que le pide un Usuario, le podemos pasar un Administrador
// Eso es porque el tipo de dato es una interfaz
// Y admin implementa a esa Interfaz. Entonces es valido.

func auth(user Person) string {
	if user.Permissions() > 4 {
		return user.Name() + " tiene permisos de administrador"
	}
	return user.Name() + " No tiene permisos de administrador"
}

type Editor struct {
	name string
}

// ----------------------------------------

func (this Editor) Permissions() int {
	return 3
}

func (this Editor) Name() string {
	return this.name
}

// ----------------------------------------

func main15() {
	admin := Admin{"Gabiel"}
	editor := Editor{"Alejandro"}

	usuarios := []Person{Admin{"Sapetin1"}, Editor{"Sapetin2"}}
	for _, usuario := range usuarios {
		fmt.Println(auth(usuario))
	}

	fmt.Println(auth(admin))
	fmt.Println(auth(editor))
}
