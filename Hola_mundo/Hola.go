package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Hello())
	var name string
	var age int

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&name)
	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&age)

	fmt.Printf("Hola, me llamo %s y tengo %d años.\n", name, age)
}
