package main

import "fmt"

// El keyword type indica que va a definirse un nuevo tipo
// El nombre es el segundo keyword, que es User
// Y el tercer dato indica que vamos a definir una estructura
type User struct {
	edad             int
	nombre, apellido string
	// Se colocan las propiedades con su tipo de dato
	// También podemos declarar por separados los campos
	// apellido string
}

func main12() {
	/*
		El struct puede ir por fuera para que sea global
	*/

	var uriel User
	fmt.Println(uriel)
	// Se define cada uno de los atributos en su valor 0
	// {0 "" ""}

	// No tienen un orden para ser asignados los atributos
	gabo := User{nombre: "Gabo", apellido: "Delgado", edad: 15}
	fmt.Println(User{nombre: "Gabriel", apellido: "Sape"})
	fmt.Println(gabo)

	// También podemos...
	alejandro := User{20, "Sape", "Delgado"}
	// Aquí si los declaramos por Orden y cada elemento debe tener un valor
	fmt.Println(alejandro)

	// También tenemos keyword new

	nuevo := new(User)
	nuevo.nombre = "Gabriel"
	// Aquí nos devuelve &{0  }
	// Es como que es un puntero, así que tenemos que utilizar ese valor
	// Lo hacemos accediendo como *nuevo y luego buscamos sus atributos
	fmt.Println(nuevo)
	fmt.Println(*nuevo)
	// Pero al parecer el lenguaje te deja acceder a los valores directamente. Jummm
	fmt.Println(nuevo.nombre)
}
