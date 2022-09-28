package main

import "fmt"

// This is a function
func HelloWorld(s string) string {
	return "Hello " + s
}

// entrypoint
func main() {
	fmt.Println(HelloWorld("there 🤖"))

	garfield := Cat{
		catFur: fur{
			color:    "orange",
			hasSpots: true,
		},
	}

	var schrodingerCat Cat

	schrodingerCat.catFur.color = "don't know"

	fmt.Println(garfield, schrodingerCat)

}

type Cat struct {
	catFur fur
}

type fur struct {
	color    string
	hasSpots bool
}
