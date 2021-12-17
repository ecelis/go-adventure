package lib

import (
	"fmt"
	"strings"
)

type location struct {
	description string
	tag         string
}

var locs = []location{
	{description: "an open field", tag: "field"},
	{description: "a little cave", tag: "cave"},
}
var playerLocation int = 0

func Execute(input string) bool {
	var token []string = strings.Split(input, " ")
	if len(token) >= 1 {
		switch token[0] {
		case "quit":
			return false
		case "look":
			if len(token) < 2 {
				return false
			}
			executeLook(token[1])
		case "go":
			if len(token) < 2 {
				return false
			}
			executeGo(token[1])
		default:
			fmt.Printf("I don't know how to do %s", token[0])
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
