package main

import "fmt"

func duplicar(n int) int {
	return n * 2
}

func triplicar(n int) int {
	return n * 3
}

func aplicar(f func(int) int, n int) int {
	return f(n)
}

func main() {
	resultado1 := aplicar(duplicar, 5)
	resultado2 := aplicar(triplicar, 5)

	fmt.Println(resultado1)
	fmt.Println(resultado2)
}
