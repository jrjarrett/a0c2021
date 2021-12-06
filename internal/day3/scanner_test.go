package day3

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_DayThreePart1(t *testing.T) {
	t.Run("test AOC example", func(t *testing.T) {
		s := Scanner{}
		readings, err := s.GetReactorInput("/Users/jarrett/src/aoc2021/testData/day3/day3test.txt")
		if err != nil {
			log.Fatal(err)
		}
		transform := s.RotateArrayCW(readings)
		gamma, epsilon := s.CalculateGammaEpsilonValue(transform)
		assert.Equal(t, 22, gamma, "Gamma value")
		assert.Equal(t, 9, epsilon, "Epsilon value")
	})

}

func Test_DayThreePart2(t *testing.T) {
	t.Run("test AOC example for O2 levels", func(t *testing.T) {
		expected := []int{1, 0, 1, 1, 1}
		s := Scanner{}
		readings, err := s.GetReactorInput("/Users/jarrett/src/aoc2021/testData/day3/day3test.txt")
		if err != nil {
			log.Fatal(err)
		}
		result := s.CalculateO2CO2Values(readings, 0, O2)
		assert.Equal(t, expected, result)
	})
	t.Run("test AOC example for CO2 levels", func(t *testing.T) {
		expected := []int{0, 1, 0, 1, 0}
		s := Scanner{}
		readings, err := s.GetReactorInput("/Users/jarrett/src/aoc2021/testData/day3/day3test.txt")
		if err != nil {
			log.Fatal(err)
		}
		result := s.CalculateO2CO2Values(readings, 0, CO2)
		assert.Equal(t, expected, result)
	})
	t.Run("Convert to decimal works", func(t *testing.T) {
		input := []int{0, 1, 0, 1, 0}
		s := Scanner{}
		assert.Equal(t, 10, s.ConvertToDecimal(input))
	})

	t.Run("Put it all together for AOC test input", func(t *testing.T) {
		s := Scanner{}
		readings, err := s.GetReactorInput("/Users/jarrett/src/aoc2021/testData/day3/day3test.txt")
		if err != nil {
			log.Fatal(err)
		}
		o2Readings := s.CalculateO2CO2Values(readings, 0, O2)
		co2Readings := s.CalculateO2CO2Values(readings, 0, CO2)
		o2 := s.ConvertToDecimal(o2Readings)
		co2 := s.ConvertToDecimal(co2Readings)
		assert.Equal(t, 230, o2*co2)

	})
}
