package day3

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
)

type Scanner struct{}

type Gases int

const (
	O2 Gases = iota
	CO2
)

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

// bitposition starts at 0 -> length of columns (in test data, that's 5; input it's 16)

func (s *Scanner) CalculateO2CO2Values(readings [][]int, bitposition int, gasType Gases) []int {
	readingCols := len(readings[0])

	if bitposition > readingCols {
		fmt.Printf("cols is %d, bitposition is %d", readingCols, bitposition)
		//panic("We've run off the end of the columns")
	}
	if len(readings) == 1 {
		return readings[0] // will need to convert from binary to decimal
	} else {
		transform := s.RotateArrayCW(readings)
		rows := len(transform[0])
		//cols := len(transform[0])
		// Look for the numbers with the nth bit set and keep. Recurse back into this routine
		oneIsSignificant := 0
		zeroIsSignificant := 0
		for i := 0; i < rows; i++ {

			if transform[bitposition][i] == 1 {
				oneIsSignificant = oneIsSignificant + 1
			} else {
				zeroIsSignificant = zeroIsSignificant + 1
			}

		}
		// we've gone thru all of the numbers, we know which bit is significant

		var filteredReadings [][]int
		switch gasType {
		case O2:
			if oneIsSignificant >= zeroIsSignificant {
				// filter and keep only the numbers whose bitposition matches
				for ii := 0; ii < rows; ii++ {
					if readings[ii][bitposition] == 1 {
						filteredReadings = append(filteredReadings, readings[ii])
					}
				}

			} else {
				// filter and keep only the numbers whose bitposition matches
				for ii := 0; ii < rows; ii++ {
					if readings[ii][bitposition] == 0 {
						filteredReadings = append(filteredReadings, readings[ii])
					}
				}
				// do something?
			}
			break
		case CO2:
			if oneIsSignificant >= zeroIsSignificant {
				// filter and keep only the numbers whose bitposition is the OPPOSITE
				for ii := 0; ii < rows; ii++ {
					if readings[ii][bitposition] == 0 {
						filteredReadings = append(filteredReadings, readings[ii])
					}
				}

			} else {
				// filter and keep only the numbers whose bitposition matches
				for ii := 0; ii < rows; ii++ {
					if readings[ii][bitposition] == 1 {
						filteredReadings = append(filteredReadings, readings[ii])
					}
				}
			}
		}

		return s.CalculateO2CO2Values(filteredReadings, bitposition+1, gasType)
	}

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
	return reactorReadings, nil

}

func (s *Scanner) ConvertToDecimal(reading []int) int {
	magnitude := len(reading)
	var value int
	for i := 0; i < magnitude; i++ {
		if reading[i] == 1 {
			value = value + int(math.Pow(2, float64(magnitude-1-i)))

		}
	}
	return value
}
