package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("Abriendo archivo...")
	defer fmt.Println("Cerrando archivo...") // Permite ejecutar una función al finalizar la función actual, incluso si ocurre un error o se produce un panic.

	fmt.Println("Leyendo archivo...")
}

func EjemploPanic() {
	defer func() {
		reco := recover() // La función recover se utiliza para manejar un panic. Se debe llamar dentro de una función defer para capturar el error generado por panic y evitar que el programa se detenga abruptamente.
		if reco != nil {
			fmt.Println("Se ha recuperado del panic:", reco)
			log.Fatalf("Ocurrió un error que generó panic %v", reco)
		}
	}()
	a := 10
	if a > 5 {
		panic("!Ocurrió un error inesperado¡") // La función panic se utiliza para generar un error de forma intencional. Detiene la ejecución normal del programa y comienza a buscar una función defer que pueda manejar el error.
	}

}

func EjemploRecover() {
	b := 20
	if b > 15 {
		recover() // La función recover se utiliza para manejar un panic. Se debe llamar dentro de una función defer para capturar el error generado por panic y evitar que el programa se detenga abruptamente.
		fmt.Println("Se ha recuperado del panic")
	}
}
