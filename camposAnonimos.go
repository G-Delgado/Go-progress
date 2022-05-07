package main

import "fmt"

type Human struct {
	name string
}

func (this Human) hablar() string {
	return "Bla bla bla"
}

type Dummy struct {
	name string
}

type Tutor struct {
	// El tutor tiene un campo humano
	Human
	Dummy
}

// Podemos sobreescribir el método que ya habíamos creado
// En la estructura de Human

func (this Tutor) hablar() string {
	return this.Human.hablar() + " Bienvenidos a sape"
	// return "Bienvenidos a Sape"
}

func main14() {
	// Esto funciona
	// Es un extends (herencia) sin utilizar palabras muy claves
	// Son campos anonimos que se heredan de otros struct
	// Cuando se hereda el mismo atributo de 2 struct diferentes
	// Nos dice que es un selector ambiguo
	// Pues no sabe si vamos por name de Dummy o de human
	// Se arregla con newTutor.Human.name ó newTutor.Dummy.name
	newTutor := Tutor{Human{"Sape"}, Dummy{"Uriel"}}
	fmt.Println(newTutor.Human.name)

	fmt.Println(newTutor.hablar())
}
