package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello para persona especifica
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("El nombre esta vacio")
	}
	message := fmt.Sprintf(RandomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func RandomFormat() string {
	formats := []string{
		"Hola, %v bienvendio",
		"Que bueno verte, %v",
		"Saludos cordiales, %v",
	}

	return formats[rand.Intn(len(formats))]
}
