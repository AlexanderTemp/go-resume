package greetings

import (
	"errors"
	"fmt"
)

func Hello2(name string) (string, error) {
	// Si no se envía nombre
	if name == "" {
		return "", errors.New("no se envío el nombre")
	}

	message := fmt.Sprintf("Hola, %v. Bienvenido!!!", name)
	return message, nil
}
