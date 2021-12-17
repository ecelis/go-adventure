package lib

// TODO learn to use pointers
type object struct {
	description string
	tag         string
	location    Location
}

var Objects = []object{
	{description: "an open field", tag: "field"},
	{description: "a little caev", tag: "cave"},
	{description: "a silver coin", tag: "silver"},
	{description: "a gold coin", tag: "gold", cave},
	{description: "a burly guard", tag: "guard", field},
	{description: "yourself", tag: "yourself", field},
}
