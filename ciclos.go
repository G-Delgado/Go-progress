package main

import "fmt"

func main6() {
	// Es basicamente igual a un ciclo en js o python
	// Sin paréntesis y se inicia i con :=
	for i := 0; i < 10; i++ {
		fmt.Printf("Soy un capo %d\n", i)
	}

	i := 0
	// Para simular un while!
	// for i < 10 {
	// 	fmt.Println(i)
	// }

	for i < 10 {
		if i == 2 {
			i++
			continue
		}
		fmt.Println(i)
		i++
		if i == 9 {
			fmt.Println("Llegué al 9!")
			break
		}
	}

}
