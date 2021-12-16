package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ecelis/go-adventure/lib"
)

var input string

func main() {
	fmt.Println("Welcome to Little Cave Adventure.")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\n--> ")

		scanner.Scan()
		input = scanner.Text()
		if len(input) != 0 {
			fmt.Println(input)
			if !lib.Execute(input) {
				return
			}
		} else {
			break
		}
	}
	fmt.Println("\nBye!")
}
