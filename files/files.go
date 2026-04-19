package files

import (
	//"bufio"
	"fmt"
	//"ioutil"
	"os"

	"github.com/farithem/godesde0/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.TablaNumerica()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Erro al crear archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.TablaNumerica()
	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar el archivo")
	}
}

func Append(fileName string, texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo" + err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error al escribir en el archivo" + err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeerArchivo() {
	archivo, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}
	fmt.Println(string(archivo))
}
