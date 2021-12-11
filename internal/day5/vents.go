package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Vents struct {}

type Point struct {
	X int
	Y int
}

type VentLine struct {
	From Point
	To Point
}




func (v *Vents) CreateVentLinesFromInput(fileName string, isPartOne bool) ([]VentLine, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ventLines := make([]VentLine,0)
	for scanner.Scan() {
		ventLine, err := createVentLine(scanner.Text())
		if err != nil {
			return nil, err
		}
		if isPartOne {
			if ventLine.From.X == ventLine.To.X || ventLine.From.Y == ventLine.To.Y {
				ventLines = append(ventLines, *ventLine)
			}
		} else {
			ventLines = append(ventLines, *ventLine)
		}
	}
	return ventLines,err
}

func (v *Vents) FindHotSpots(ventLines []VentLine) map[Point]int {
	hotSpots := make(map[Point]int)
	for _, ventLine := range ventLines {
		var start, finish int
		switch {
		case ventLine.From.X == ventLine.To.X:
			// fmt.Printf("The X values match for %v\n", ventLine)
			if ventLine.From.Y > ventLine.To.Y {
				start = ventLine.To.Y
				finish = ventLine.From.Y
			} else {
				start = ventLine.From.Y
				finish = ventLine.To.Y
			}
			for i :=start; i <= finish; i++ {
				p, exists :=hotSpots[Point{X:ventLine.From.X, Y:i}]
				if (exists) {
					hotSpots[Point{X:ventLine.From.X, Y:i}] = p +1
				} else {
					hotSpots[Point{X:ventLine.From.X, Y:i}] = 1
				}
			}
			break
		case ventLine.From.Y == ventLine.To.Y:
			// fmt.Printf("The Y values match for %v\n", ventLine)
			if ventLine.From.X > ventLine.To.X {
				start = ventLine.To.X
				finish = ventLine.From.X
			} else {
				start = ventLine.From.X
				finish = ventLine.To.X
			}
			for i :=start; i <= finish; i++ {
				p, exists :=hotSpots[Point{X:i, Y:ventLine.From.Y}]
				if (exists) {
					hotSpots[Point{X:i, Y:ventLine.From.Y}] = p +1
				} else {
					hotSpots[Point{X:i, Y:ventLine.From.Y}] = 1
				}
			}
			break
		default:
			fmt.Printf("This line is not processed %v\n", ventLine)
		}
	}
	return hotSpots
}

func (v * Vents) CalculateDayOneAnswer(hotSpots map[Point]int) int {
	value := 0
	for _, spot := range hotSpots {
		if spot > 1 {
			value++
		}
	}
	return value
}
func createVentLine(input string) (*VentLine, error) {
	components := strings.Fields(input)

	// components is now [ "x1,y1", "->", "x2,y2"]

	fromPoint, err := createPoint(components[0]) // ["x1,y1"]
	if err != nil {
		return nil, err
	}
	toPoint, err := createPoint(components[2]) // ["x2,y2"]
	if err != nil {
		return nil, err
	}

	return &VentLine{From: *fromPoint, To: *toPoint}, nil

}

func createPoint(input string) (*Point, error) {
	pointAsString := strings.Split(input, ",")
	pointX, err := strconv.Atoi(pointAsString[0])
	if err != nil {
		return nil, err
	}
	pointY, err := strconv.Atoi(pointAsString[1])
	if err != nil {
		return nil, err
	}
	return &Point{X: pointX, Y: pointY,}, nil


}
