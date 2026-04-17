package mouredev

import "fmt"

func Prueba() {
	fmt.Println("Esta es mi función de prueba")

	var nombre string = "Farithem"
	fmt.Println("El nombre es: ", nombre)

	var edad int = 25
	fmt.Println("La edad es: ", edad)

	var altura float64 = 1.75
	fmt.Println("La altura es: ", altura)

	var soltero bool = false
	fmt.Println("¿Es soltero?: ", soltero)

	// Constantes
	const pi = 3.14159
	fmt.Println("El valor de pi es: ", pi)

	// Operadores
	fmt.Println("La edad mas 1 es: ", edad+1)
	fmt.Println("La edad menos 1 es: ", edad-1)
	fmt.Println("La edad multiplicada por 2 es: ", edad*2)
	fmt.Println(edad / 2)
	fmt.Println(edad % 2)
}
