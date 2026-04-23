package funciones

import "fmt"

func Calculos() {
	var numero3, numero4 int = 10, 20

	suma := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}

	fmt.Println(suma(5, 10))

	suma = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero3 * numero4
	}
	fmt.Println(suma(5, 10))
}
