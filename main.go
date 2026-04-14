package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(1584)
	fmt.Println("Estado:", estado)
	fmt.Println("Texto:", texto)
	*/
	if os := runtime.GOOS; os == "linux" || os == "darwin" {
		fmt.Println("Estas usando Linux")
		procesador := runtime.GOARCH
		fmt.Println("Distribución:", procesador)
	} else {
		fmt.Println("Estás usando Windows")
		fmt.Println(os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Estas usando Linux")
	case "darwin":
		fmt.Println("Estas usando MacOS")
	default:
		fmt.Printf("%s no es Linux ni MacOS", os)
	}
}
