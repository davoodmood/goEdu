package main

import (
	"fmt"

	"log"

	"github.com/davoodmood/goEdu/tree/main/greeting"

	"rsc.io/quote"

	"math/rand"
	"time"
)

var messages map[string]string
var err error

func main() {

	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Han", "Lea", "Luke"}

	// Get a greeting message and print it.
	messages, err = greeting.Hellos(names) // + " " + quote.Opt()

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(messages)
	for i, message := range messages {
		fmt.Println("To " + i + ": " + message + " " + randomFormat() + "")
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		quote.Opt(),
		quote.Glass(),
		quote.Go(),
	}
	return formats[rand.Intn(len(formats))]
}
