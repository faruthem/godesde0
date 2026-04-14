package main

import (
	"fmt"

	"github.com/farithem/godesde0/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(1584)
	fmt.Println("Estado:", estado)
	fmt.Println("Texto:", texto)
}
