package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	paises["Colombia"] = "Bogotá"
	paises["Perú"] = "Lima"
	paises["Argentina"] = "Buenos Aires"
	paises["Chile"] = "Santiago"
	fmt.Println(paises)

	campeonato := map[string]int{
		"Real Madrid":   14,
		"AC Milan":      7,
		"Liverpool":     6,
		"Bayern Munich": 6,
	}
	fmt.Println(campeonato)

	for equipo, titulos := range campeonato {
		fmt.Printf("El equipo %s tiene %d títulos\n", equipo, titulos)
	}

	delete(campeonato, "Real Madrid")

	for equipo, titulos := range campeonato {
		fmt.Printf("El equipo %s tiene %d títulos\n", equipo, titulos)
	}
	puntaje, existe := campeonato["Liverpool"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)
}
