/*
 1. Crear un archivo GO en el paquete ejercicios llamado ejercicio02.go
 2. Crear una función Pública, que pida por teclado un número, valide si hay error o no(y si hay error vuelva a
    pedirlo)
 3. En base al número recibido crear la tabla numérica del mismo. Del 1 al 10 y la muestre en pantalla.
 4. En Main, llama a dicha función.
 5. Grabar, ejecutar y luego subir a tu repositorio de GitHub.
*/
package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func TablaNumerica() string {
	//texto = "" // Limpiar la variable antes de usarla
	for {
		fmt.Println("Ingrese el número: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			for i := 1; i <= 10; i++ {
				resultado := numero * i
				texto += fmt.Sprintf("%d x %d = %d\n", numero, i, resultado)
			}
			if err == nil {
				break
			}
		}
	}
	if err != nil {
		panic("El dato ingresado es incorrecto" + err.Error())
	}
	return texto
}
