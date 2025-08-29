package main

import (
	"fmt"
	"log"

	"example_2.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags((0))
	message := greetings.Hello("alex")

	fmt.Println(message)

	mensaje, err := greetings.Hello2("")

	if err != nil {
		log.Println(err)
	}
	fmt.Println(mensaje)

	mensaje2, err := greetings.Hello3("Jaime del rio")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(mensaje2)
}
