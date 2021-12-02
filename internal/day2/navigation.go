package day2

import (
	"errors"
)

type NavigationSystem struct {
	Aim int
}

type Navigation struct {
	Direction string
	Position  int
}

type Position struct {
	Horizontal int
	Depth      int
}

func (n *NavigationSystem) CalculatePositionV1(path []Navigation) (Position, error) {

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
func (n *NavigationSystem) CalculatePositionV2(path []Navigation) (Position, error) {

	var pos Position
	for _, step := range path {
		switch step.Direction {
		case "forward":
			pos.Horizontal += step.Position
			pos.Depth += n.Aim * step.Position
		case "down":
			n.Aim += step.Position

		case "up":
			n.Aim -= step.Position

		default:
			return Position{}, errors.New("Bad direction in input")
		}
	}
	return pos, nil
}
