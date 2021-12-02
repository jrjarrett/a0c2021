package day2

import "errors"

type NavigationSystem struct{}

type Navigation struct {
	Direction string
	Position  int
}

type Position struct {
	Horizontal int
	Depth      int
}

func (n *NavigationSystem) CalculatePosition(path []Navigation) (Position, error) {

	var pos Position
	for _, step := range path {
		switch step.Direction {
		case "forward":
			pos.Horizontal += step.Position
		case "down":
			pos.Depth += step.Position
		case "up":
			pos.Depth -= step.Position
		default:
			return Position{}, errors.New("Bad direction in input")
		}
	}
	return pos, nil
}
