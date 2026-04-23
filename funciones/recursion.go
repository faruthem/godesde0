package funciones

import "fmt"

func Exponencia(valor int) {
	if valor > 10000000000 {
		return
	}
	fmt.Println("Valor:", valor)
	Exponencia(valor * 2)
}
