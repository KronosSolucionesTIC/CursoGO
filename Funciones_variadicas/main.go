package main

import "fmt"

func suma(name string, nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}

	fmt.Printf("Hola %s, la suma es: %d\n", name, total)

	return total
}

func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("%T - %v\n", dato, dato)
	}
}

func main() {
	fmt.Println(suma("Diego", 10, 20, 30, 40, 50))
	fmt.Println(suma("Alejandro", 1, 2, 3, 4))
	imprimirDatos("Hola", 15, true, 3.1416)
}
