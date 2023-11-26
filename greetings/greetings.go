package greetings

import "fmt"
import "errors"

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	if name != "" {
		return message, nil
	} else {
		return "", errors.New("xxxx")
	}
}

func main() {
	Hello("Every Body")
}
