package main

import (
	"fmt"

	"github.com/ecelis/go-adventure/lib"
)

var input string

func getInput() bool {
	fmt.Printf("\n--> ")
	_, err := fmt.Scanf("%s", &input)
	return err == nil
}

func main() {
	fmt.Println("Welcome to Little Cave Adventure.")
	for getInput() {
		if !lib.Execute(input) {
			return
		}
	}
	fmt.Println("\nBye!")
}
