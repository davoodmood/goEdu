package main

import (
	"fmt"

	"log"

	"github.com/davoodmood/goEdu/greeting"

	"rsc.io/quote"
)

var message string
var err error

func main() {

	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	message, err = greeting.Hello("David") // + " " + quote.Opt()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message + " " + quote.Opt())
}
