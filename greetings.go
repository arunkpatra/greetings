package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person. An error is returned if an empty name is provided to this
// function.
func Hello(name string) (string, error) {
	// If no name was provided, return an error.
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
