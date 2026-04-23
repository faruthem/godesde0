package arreglos_slides

import "fmt"

var tabla [10]int
var matriz [20][30]int

func MuestroArreglos() {
	tabla[0] = 1
	tabla[1] = 2
	tabla[2] = 3
	tabla[3] = 4
	tabla[4] = 5
	tabla[5] = 6
	tabla[6] = 7
	tabla[7] = 8
	tabla[8] = 9
	tabla[9] = 10
	fmt.Println("Tabla:", tabla)
	for i := 0; i < len(tabla); i++ {
		fmt.Printf("Posición %d: %d\n", i, tabla[i])
	}

	matriz[0][0] = 1
	matriz[0][1] = 2
	matriz[0][2] = 3
	matriz[0][3] = 4
	matriz[0][4] = 5
	matriz[0][5] = 6
	matriz[0][6] = 7
	matriz[0][7] = 8
	matriz[0][8] = 9
	matriz[0][9] = 10
	fmt.Println("Matriz:", matriz)

}
