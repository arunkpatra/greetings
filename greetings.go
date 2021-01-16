package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person. An error is returned if an empty name is provided to this
// function.
func Hello(name string) (string, error) {
	// If no name was provided, return an error.
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map of greeting messages. The map contains a message for each of the names passed to this
// function.
func Hellos(names []string) (map[string]string, error) {
	// Create the map
	messages := make(map[string]string)
	// Iterate through the names provided and create a message for each
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns a random format from a fixed set of message formats.
func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format
	return formats[rand.Intn(len(formats))]
}
