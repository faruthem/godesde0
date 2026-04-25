package ejer_interfaces

import (
	"fmt"

	"github.com/farithem/godesde0/interfaces"
)

func HumanosRespirando(humanos interfaces.Humanos) {
	humanos.Respirar()
	fmt.Printf("Soy un/a humano %s, y estoy respirando \n", humanos.Sexo())
}

func SerVivo(vegetal interfaces.Servivo) {
	vegetal.EstaVivo()
	fmt.Printf("soy un vegetal y estoy vivo: %t \n", vegetal.EstaVivo())
	fmt.Printf("Mi clasificación vegetal es: %s \n", vegetal.ClasificacionVegetal())
}
