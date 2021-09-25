// aim is to return one of the several predefined greeting message.

package slice_greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func SliceGreetings(name string) (string, error) {
	// if no name was given,return a error with a message

	if name == "" {
		return name, errors.New("empty name")
	}

	// create a message using a random format
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

//
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// while declaring slice you omit its size in brakets like []string, tells go that the size of the array can be dynamically changed
	// Add an init function to seed the rand package with the current time.
	// Go executes init functions automatically at program startup, after global variables have been initialized.
	formats := []string{
		"Hello, %v. Welcome",
		"Great to see you, %v!",
		"Hail, %v Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
