package main

import (
	"library/animal"
)

func main() {
	// myBook := book.NewBook("Moby Dick", "Herman Melville", 300)
	// myTextBook := book.NewTextBook("Comunicacion", "Jaime Gamarra", 261,
	// 	"Santillana SAS", "Secundaria")

	// book.Print(myBook)
	// book.Print(myTextBook)
	animales := []animal.Animal{
		&animal.Perro{Nombre: "Max"},
		&animal.Gato{Nombre: "Tom"},
		&animal.Perro{Nombre: "Lupe"},
		&animal.Gato{Nombre: "Samy"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
