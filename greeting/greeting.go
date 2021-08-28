package greeting

import (
	"errors"
	"fmt"
)

var message string

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("You must enter a name as argument.")
	}

	// Return a greeting that embeds the name in a message.
	message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
