package day7

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "math"
    "testing"
)

func TestDay1(t *testing.T) {
    t.Run("Find Mode", func(t *testing.T) {
        crabPositions := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
        crabMap := make(map[int]int)
        for i := 0; i < len(crabPositions); i++ {
            val, exists := crabMap[crabPositions[i]]
            if exists {
                crabMap[crabPositions[i]] = val + 1
            } else {
                crabMap[crabPositions[i]] = 1
            }
        }
        fmt.Printf("Crabmap is %v\n", crabMap)
        var largest, largestCount int
        for position, count := range crabMap {
            if count > largestCount {
                largestCount = count
                largest = position
            }
        }
        fmt.Printf("The best position is %d\n", largest)
        var fuelUsed int
        for i := 0; i < len(crabPositions); i++ {
            fuelUsed = fuelUsed + int(math.Abs(float64(crabPositions[i]-largest)))
        }
        fmt.Printf("Total fuel used is %d\n", fuelUsed)
    })
    t.Run("day 1 test", func(t *testing.T) {
        c := CrabSub{}
        crabPositions := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
        answer := c.FuelSpendToAlign(crabPositions)
        assert.Equal(t, 37, answer)

    })

    t.Run("Test loading from file", func(t *testing.T) {
        c := CrabSub{}
        expected := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
        result := c.GetPositionsFromFile("/Users/jarrett/src/dev/golang/aoc2021/testData/day7/day7test.txt")
        assert.Equal(t, expected, result)

    })
}

func TestDay2(t *testing.T) {
    t.Run("Test day 2", func(t *testing.T) {
        c := CrabSub{}
        crabPositions := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
        result := c.FuelSpendToAlignD2(crabPositions)
        assert.Equal(t, 168, result)

    })

}
