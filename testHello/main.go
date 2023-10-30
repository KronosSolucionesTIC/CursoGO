package main

import (
	"fmt"

	"github.com/KronosSolucionesTIC/greetings"
)

func main() {
	message, err := greetings.Hello("Diego")

	if err != nil {
		fmt.Println("Ocurrio un error:", err)
		return
	}

	fmt.Println(message)
}
