package teclado

import (
	"bufio" // Importamos el paquete bufio para leer la entrada del usuario
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	fmt.Println("Ingrese número 1: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
	}
	if err != nil {
		panic("El dato ingresado es incorrecto" + err.Error())
	}
	fmt.Println("Ingrese número 2: ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
	}
	if err != nil {
		panic("El dato ingresado es incorrecto" + err.Error())
	}
	fmt.Println("Ingrese una leyenda: ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}
	fmt.Printf("Número 1: %d, Número 2: %d, Leyenda: %s\n", numero1, numero2, leyenda)
	fmt.Println(leyenda, ": ", numero1*numero2)
}
