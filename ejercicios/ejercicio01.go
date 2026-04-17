package ejercicios

import "strconv"

func Texto(valor string) (int, string) {
	//Parámetro string debe ser convertido a entero
	numero, err := strconv.Atoi(valor)
	if err != nil {
		// Manejar el error de conversión
		return 0, "Error al convertir a entero"
	}
	//Si el valor es mayor a 100, el string retornado debe decir "Es mayor a 100"
	//De los contrario, devolver el mensaje "Es menor a 100"
	if numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "Es menor a 100"
	}
}
