package main

// you declare a main package (a package is a way to group functions, and its made up of all the files same directory)
//a main function executes by default when you run the main package.

import (
	"fmt"
	"log"

	"multi-greetings.com/multi_greetings"
	"ravi-greeting.com/greetings"
	"slice-greetings.com/slice_greetings"
)

func main() {
	// set properties of the predefined logger
	log.SetPrefix("slice greetings: ")
	log.SetFlags(0)

	// Request a greeting message
	message, err := slice_greetings.SliceGreetings("Ravindra")
	if err != nil {
		log.Fatal(err)
	}
	multiGreetings()
	raviGreetings()
	fmt.Println(message)
}

func raviGreetings() {
	log.SetPrefix("raviGreetings --- ")
	log.SetFlags(0)

	message, err := greetings.Hello("Ravindra")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}

func multiGreetings() {
	log.SetPrefix("Multi Greetings: ")
	log.SetFlags(0)
	messages, err := multi_greetings.MultiGreetigns([]string{"Ravindra", "Paladugu"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
