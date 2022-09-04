package greeting

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var messages map[string]string = make(map[string]string)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("you must enter a name as argument")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprint(randomFormat()) // an intentional buggy line to see error in the test.
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}

	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
