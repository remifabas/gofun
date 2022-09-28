package gofun

import "fmt"

// Capitalized -> public
type Cat struct {
	catFur fur
}

// private
type fur struct {
	color    string
	hasSpots bool
}

func NewCat(color string, spots bool) Cat {
	return Cat{
		catFur: fur{
			color:    color,
			hasSpots: spots,
		},
	}
}

//receiver function
func (c *Cat) Color() string {
	return c.catFur.color
}

type Prey struct {
	isYummy bool
}

type Hunter interface {
	Hunt() (Prey, error)
}

func (c *Cat) Hunt() (Prey, error) {
	fmt.Println("Cat hunt implementation")
	return Prey{isYummy: true}, nil
}
