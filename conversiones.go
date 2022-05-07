package main

// Tampoco podemos mantener paquetes importados sin usarlos
import (
	"fmt"
	"strconv"
)

func main3() {
	edad := 22
	edad2 := "22"
	// Da error pues es str + int
	// fmt.Println("Mi edad es " + edad)

	edad_str := strconv.Itoa(edad)
	// La función Atoi retorna dos valores, por tanto no se puede inicializar así nada más
	// edad_int := strconv.Atoi(edad2)

	// De esta forma podemos recibir los 2 valores
	// edad_int, err := strconv.Atoi(edad2)

	// Ponemos guion bajo para avisarle que no nos importa el segundo valor
	edad_int, _ := strconv.Atoi(edad2)

	fmt.Println("Mi edad es " + edad_str)
	fmt.Println(edad_int + 10)
}
