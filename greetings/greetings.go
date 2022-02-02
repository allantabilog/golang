package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {

	if name == "" {
		return name, errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// create a map to associate names with messages
	messages := make(map[string]string)

	// loop through the received slice
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// associate the retrieved name message with the name
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
		"Hallhallo %v fancy meetin ya!!",
	}

	return formats[rand.Intn(len(formats))]
}
