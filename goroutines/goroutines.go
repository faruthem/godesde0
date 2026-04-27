package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLentooo(nombre string, ch chan bool) {
	multiplo := 0
	i := 1
	letras := strings.Split(nombre, "") // La función strings.Split se utiliza para dividir una cadena en partes más pequeñas, utilizando un separador específico. En este caso, se está dividiendo la cadena "nombre" en letras individuales, utilizando una cadena vacía como separador.
	for _, letra := range letras {
		multiplo += multiplo
		fmt.Println(multiplo * 1)
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)

		fmt.Println(strings.Repeat("*", i))
		fmt.Println("Contador:", i)
		i++
	}
	ch <- true
}
