package arreglos_slides

import "fmt"

var tablas []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var arreglo [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func MuestroSlides() {
	fmt.Println("Muestro los arreglos")
	fmt.Println("Tablas:", tablas)
	porcion := tablas[2:5]
	fmt.Println("Porción:", porcion)
	fmt.Println("Arreglo:", arreglo)

}

func Capacidad() {
	elementos := make([]int, 5, 10)
	fmt.Printf("Largo %d, Capacidad %d \n", len(elementos), cap(elementos))
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
		fmt.Printf("\nLargo %d, Capacidad %d \n", len(nums), cap(nums))
	}

}
