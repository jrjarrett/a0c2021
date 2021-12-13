package day7

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CrabSub struct{}

func (c *CrabSub) GetPositionsFromFile(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	positions := make([]int, 0)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	vals := strings.Split(scanner.Text(), ",")
	for i := 0; i < len(vals); i++ {
		crab, err := strconv.Atoi(vals[i])
		if err != nil {
			fmt.Printf("Bad input at position %d\n", i)
			panic("Bad input from file")
		}
		positions = append(positions, (crab))
	}
	fmt.Printf("Number of inputs is %d\n", len(positions))
	return positions
}

func (c *CrabSub) FuelSpendToAlign(crabPositions []int) int {
	sort.Ints(crabPositions)
	var median int
	medianPosition := len(crabPositions) / 2
	if medianPosition%2 == 0 {
		median = (crabPositions[medianPosition-1] + crabPositions[medianPosition]) / 2
	} else {
		median = crabPositions[medianPosition]
	}
	var medianFuelUsed int
	for i := 0; i < len(crabPositions); i++ {
		medianFuelUsed = medianFuelUsed + int(math.Abs(float64(crabPositions[i]-median)))
	}

	fmt.Printf("Median fuel used is %d\n", medianFuelUsed)
	return medianFuelUsed

}

func (c *CrabSub) FuelSpendToAlignD2(crabPositions []int) int {

	minFuelUsed := -1
	for i := 0; i < len(crabPositions); i++ {
		fuelUsed := 0
		for j := 0; j < len(crabPositions); j++ {
			positionsToMove := int(math.Abs(float64(i - crabPositions[j])))
			fuelUsed = fuelUsed + (positionsToMove*(positionsToMove+1))/2
		}
		if minFuelUsed == -1 {
			minFuelUsed = fuelUsed
		}

		if fuelUsed < minFuelUsed {
			minFuelUsed = fuelUsed
		}
		//		fmt.Printf("Fuel used to move to %d is %d\n", i, fuelUsed)
	}

	fmt.Printf("Min fuel used is %d\n", minFuelUsed)
	return minFuelUsed

}
