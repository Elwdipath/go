package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set predefined logger properties
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//slice of names
	names := []string{"Jordan", "Fry", "Bender"}
	// Request greeting message
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	for _, message := range messages {
		fmt.Println(message)
	}
}
