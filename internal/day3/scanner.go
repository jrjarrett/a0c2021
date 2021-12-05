package day3

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
)

type Scanner struct{}

func (s *Scanner) RotateArrayCW(orig [][]int) [][]int {
	rows := len(orig)
	cols := len(orig[0])

	neo := make([][]int, cols)
	for i := range neo {
		neo[i] = make([]int, rows)
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			neo[c][r] = orig[r][c]

		}
	}

	return neo
}

func (s *Scanner) CalculateGammaEpsilonValue(readings [][]int) (int, int) {
	rows := len(readings)
	cols := len(readings[0])
	var gammaValue int
	var epsilonValue int

	for i := 0; i < rows; i++ {
		oneIsSignificant := 0
		zeroIsSignificant := 0
		for j := 0; j < cols; j++ {
			if readings[i][j] == 1 {
				oneIsSignificant = oneIsSignificant + 1

			} else {
				zeroIsSignificant = zeroIsSignificant + 1
			}

		}
		if oneIsSignificant > zeroIsSignificant {
			gammaValue += int(math.Pow(2, (float64)(rows-1-i)))
		} else {
			epsilonValue += int(math.Pow(2, (float64)(rows-1-i)))
		}

	}
	return gammaValue, epsilonValue
}

func (s *Scanner) GetReactorInput(fileName string) ([][]int, error) {

	var reactorReadings [][]int
	file, err := os.Open(fileName)
	if err != nil {
		return reactorReadings, err
	}
	lineCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount = lineCount + 1
		var binaryNumber []int
		line := scanner.Text()
		for _, c := range line {
			if c == '1' {
				binaryNumber = append(binaryNumber, 1)
			} else if c == '0' {
				binaryNumber = append(binaryNumber, 0)
			} else {
				return nil, errors.New("Illegal input in line " + line)
			}
		}

		reactorReadings = append(reactorReadings, binaryNumber)
	}
	err = file.Close()
	if err != nil {
		return reactorReadings, err
	}
	if scanner.Err() != nil {
		return reactorReadings, scanner.Err()
	}
	fmt.Printf("Line count is %d\n", lineCount)
	return reactorReadings, nil

}
