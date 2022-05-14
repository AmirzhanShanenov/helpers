package helper

import "log"

// ErrCheck Switch case log fatal,println, and others
func ErrCheck(err error, message string, level string) {

	if err != nil {
		switch level {
		case "":
			log.Println(message, err)
		case "Info":
			log.Println("Info", message, err)

		case "Fatal":
			log.Fatal("Fatal", message, err)
		case "Panic":
			log.Panic("Panic", message, err)
			//case "Trace":
			//case "Debug":
			//case "Warn":
			//case "Error":
		}
	}
}
