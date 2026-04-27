package main

import "github.com/farithem/godesde0/defer_panic"

func main() {
	/*variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(1584)
	fmt.Println("Estado:", estado)
	fmt.Println("Texto:", texto)

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
	mouredev.Prueba()
	textoPrueba := "45"
	numero, mensaje := ejercicios.Texto(textoPrueba)
	fmt.Println("Número:", numero)
	fmt.Println("Mensaje:", mensaje)

	teclado.IngresoNumeros()

	iteraciones.Iterar()
	*/
	//fmt.Println(ejercicios.TablaNumerica())
	//files.GrabaTabla()
	//files.SumaTabla()
	//files.LeerArchivo()
	//funciones.Calculos()
	//funciones.LlamarClosure()
	//funciones.Exponencia(2)
	//arreglos_slides.MuestroArreglos()
	//arreglos_slides.Capacidad()
	//mapas.MostrarMapas()
	//users.AltaUsuario()
	//Kevin := new(modelos.Hombre)
	//ejer_interfaces.HumanosRespirando(Kevin)
	//María := new(modelos.Mujer)
	//ejer_interfaces.HumanosRespirando(María)
	//Lechuga := new(modelos.Vegetal)
	//ejer_interfaces.SerVivo(Lechuga)
	defer_panic.VemosDefer()
	defer_panic.EjemploPanic()
}
