package main

import (
	"net/http"
)

// Para evitar estos problemas, lo que hacemos es crerar una carpeta llamada public
/*
	En esta carpeta estarán todos los archivos que queremos que puedan ver los usuarios
*/

func main() {
	fileServer := http.FileServer(http.Dir("public"))

	// http.Handle("/", http.StripPrefix("/", fileServer))
	// En este caso estaríamos buscando a través de la carpeta public
	// http.Handle("/public/", http.StripPrefix("/", fileServer))
	http.Handle("/public/", http.StripPrefix("/public", fileServer))

	http.ListenAndServe(":8000", nil)
}
