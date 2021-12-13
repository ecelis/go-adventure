package lib

import (
	"fmt"
	"strings"
)

func Execute(input string) bool {
	var verb []string = strings.Split(input, " ")
	// var noun = "\n"
	fmt.Println(verb)
	if len(verb) > 0 {
		switch verb[0] {
		case "quit":
			return false
		case "look":
			fmt.Println("It is very dark in here.")
		case "go":
			fmt.Println("It's too dark to go anywhere.")
		default:
			fmt.Printf("I don't know how to %s", verb[0])
		}
	}
	return true
}
