package lib

import (
	"fmt"
	"strings"
)

type location struct {
	description string
	tag         string
}

/*
	locs := []location{
		location{description: "an open field", tag: "field"},
		location{description: "a little cave", tag: "cave"},
	}

*/
var locs = make([]location, 2)
var playerLocation int = 0

func Execute(input string) bool {
	var token []string = strings.Split(input, " ")
	// var noun = "\n"
	fmt.Println(token)
	if len(token) > 0 {
		switch token[0] {
		case "quit":
			return false
		case "look":
			// fmt.Println("It is very dark in here.")
			executeLook(token[1])
		case "go":
			// fmt.Println("It's too dark to go anywhere.")
			executeGo(token[1])
		default:
			fmt.Printf("I don't know how to %s", token[0])
		}
	}
	return true
}

func executeLook(noun string) {
	switch noun {
	case "around":
		fmt.Printf("Your are in %s.\n", locs[playerLocation].description)
	default:
		fmt.Println("I don't undestand what you want to see.")
	}
}

func executeGo(noun string) {
	for i := 0; i < len(locs); i++ {
		if noun != "" && noun == locs[i].tag {
			switch i {
			case playerLocation:
				fmt.Println("You can't get much closer than this.")
			default:
				fmt.Println("OK")
				playerLocation = i
				executeLook("around")
			}
			return
		}
	}
	fmt.Println("I don't understand where you want to go.")
}
