package dummy

// Este no  va a funcionar por fuera del paquete pues al tener la primera letra en mayuscula, lo toma como privado y no público
// Si usamos el atributo con una mayuscula en camel case no la toma como público  y la exporta y podemos hacer dummy.HolaMundo
// Cosa que con hola_mundo no se puede
var hola_mundo string

// Init se llama justo al importar el paquete
func init() {
	hola_mundo = "Hola :3"
}

func holaMundoDos() string {
	return "Hola Mundo 2!"
}

func HolaMundo() string {
	return "Hola Mundo!"
}
