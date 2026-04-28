package middleware

import "fmt"

func Sumar(a, b int) int {
	return a + b
}

func Restar(a, b int) int {
	return a - b
}

func Multiplicar(a, b int) int {
	return a * b
}

func Dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("no se puede dividir por cero")
	}
	return a / b, nil
}

func MiMiddleware() {
	fmt.Println("Ejecutando middleware...")

	result := operacionesMidd(Sumar)(2, 3)
	fmt.Println("Resultado de la operación:", result)
	result = operacionesMidd(Restar)(2, 3)
	fmt.Println("Resultado de la operación:", result)
	result = operacionesMidd(Multiplicar)(2, 3)
	fmt.Println("Resultado de la operación:", result)
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operación")
		return f(a, b)
	}
}
