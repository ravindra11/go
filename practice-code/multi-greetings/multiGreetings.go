package multi_greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// slice of names as input
// multi greeting messages with provided names
/* -------------------------------------------------------------------------- */
// split the requirement
// 1. greeting of each individual name, need a function whos return value can be greeting with name

// loop the provided input names , so that each of name can be send to the 1st function to get greeting of that particular name

// add the catched greeting into map
// return map

func MultiGreetigns(names []string) (map[string]string, error) {

	// check len of input names
	// if empty, return empty map and err
	if len(names) == 0 {
		return make(map[string]string), errors.New("empty name")
	}

	// a map to associate names with messages
	messages := make(map[string]string)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.

	for _, name := range names {
		message, err := getMessage(name)
		if err != nil {
			return nil, err
		}
		// in the map, associate the retrived message with the name
		messages[name] = message
	}
	return messages, nil
}

func getMessage(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf(greeting(), name), nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// return random greeting message
func greeting() string {
	greetings := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return greetings[rand.Intn(len(greetings))]
}
