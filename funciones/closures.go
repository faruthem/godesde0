package funciones

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0

	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func LlamarClosure() {
	tabladel := 2

	MiTabla := Tabla(tabladel)
	OtraTabla := Tabla(3)

	for range 10 {
		println(MiTabla())
	}
	for range 10 {
		println(OtraTabla())
	}
}
