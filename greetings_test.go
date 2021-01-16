package greetings

import "testing"
import "regexp"

// TestHelloWithValidName calls greetings.Hello with a valid name and checks for a valid return message
func TestHelloWithValidName(t *testing.T) {
	name := "Arun"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Arun")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Arun") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloWithEmptyName calls greetings.Hello with an empty name and looks for an error to be returned.
func TestHelloWithEmptyName(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

// TestHellosWithValidNames calls greetings.Hellos with a valid list of names and checks that the responses are correct.
func TestHellosWithValidNames(t *testing.T) {
	names := []string{"Arun", "Tim", "John"}

	// Get back messages
	messages, err := Hellos(names)

	// We should not get an error
	if err != nil {
		t.Fatalf("Hellos(Arun, Tim, John) returned an error %v", err)
	}

	// Iterate through the names and check against results
	for _, name := range names {
		// First we must find for a result
		want := regexp.MustCompile(`\b` + name + `\b`)
		msg, found := messages[name]
		if found == false {
			t.Fatalf("Message not found for %v", name)
		}
		if !want.MatchString(msg) || err != nil {
			t.Fatalf("Hello(%v) = %q, %v, want match for %#q, nil", name, msg, err, want)
		}
	}
}

func TestHellosWithInvalidName(t *testing.T) {
	names := []string{"Arun", "", "John"}
	// Get back messages
	_, err := Hellos(names)
	// We should not get an error
	if err == nil {
		t.Fatalf("Hellos(Arun, \"\", John) returned no error %v", err)
	}
}
