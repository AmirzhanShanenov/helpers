package helper

import "log"

// ErrCheck Switch case log fatal,println, and others
func ErrCheck(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}
}
