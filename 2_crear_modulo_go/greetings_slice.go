package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello3(name string) (string, error) {
	if name == "" {
		return name, errors.New("nombre vacío")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hola, %v Bienvenido!",
		"Un gusto verte %v",
		"Que bueno que volviste %v",
	}

	return formats[rand.Intn(len(formats))]
}
