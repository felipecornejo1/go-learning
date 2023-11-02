package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings error: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Felipe")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
