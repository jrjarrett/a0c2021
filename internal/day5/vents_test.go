package day5

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestParsing(t *testing.T) {
	t.Run("Fields", func(t *testing.T) {
		s := "0,9 -> 5,9"
		components := strings.Fields(s)
		frompoint := strings.Split(components[0], ",")
		fmt.Printf("Fields are %q\n",components)
		fmt.Printf("from point are %q\n",frompoint)
		fromint, err := strconv.Atoi(frompoint[0])
		if err != nil {
			fmt.Printf("\n\nerr is %#v\n", err)
		}
		fmt.Printf("from value is %d\n", fromint)

	})
}

func TestVentLine(t *testing.T) {
	t.Run("Test Creating a Vent Line", func(t *testing.T) {
		expected := VentLine{
			From: Point{X: 0, Y:9},
			To:   Point{X:5, Y:9},
		}
		s := "0,9 -> 5,9"
		result ,err := createVentLine(s)
		assert.Nil(t, err)
		assert.Equal(t, &expected, result)
	})
	t.Run("Test Vent Lines from AOC test input", func(t *testing.T) {

		v := Vents{}
		result ,err := v.CreateVentLinesFromInput("/Users/jarrett/src/dev/golang/aoc2021/testData/day5/day5test.txt", true)
		assert.Nil(t, err)
		assert.Equal(t, 6, len(result))
	})

	t.Run("Test Horizontal Line", func(t *testing.T) {
		v := Vents{}
		expected := map[Point]int{Point{X:0, Y:9}:2, Point{X:1, Y:9}:2, Point{X:2, Y:9}:2, Point{X:3, Y:9}:1, Point{X:4, Y:9}:1, Point{X:5, Y:9}:1}

		ventLines := make([]VentLine,0)
		testReadings :=  []string{"0,9 -> 5,9", "0,9 -> 2,9"}
		for _,s := range testReadings {
			ventLine, err := createVentLine(s)
			assert.Nil(t, err)
			ventLines = append(ventLines, *ventLine)
		}
		fmt.Printf("ventLines is %v\n", ventLines)
		hotSpots := v.FindHotSpots(ventLines)
		fmt.Printf("hotspots is %#v\n",hotSpots)
		assert.Equal(t, expected, hotSpots)
	})
	t.Run("Test Horizontal Line Intersect", func(t *testing.T) {
		v := Vents{}
		ventLines := make([]VentLine, 0)
		testReadings := []string{"0,9 -> 5,9", "0,9 -> 2,9"}
		for _, s := range testReadings {
			ventLine, err := createVentLine(s)
			assert.Nil(t, err)
			ventLines = append(ventLines, *ventLine)
		}
		fmt.Printf("ventLines is %v\n", ventLines)
		hotSpots := v.FindHotSpots(ventLines)
		fmt.Println(hotSpots)
		result := v.CalculateAnswer(hotSpots)
		assert.Equal(t, 3, result)
	})
	t.Run("Test Vertical Line", func(t *testing.T) {
		expected := map[Point]int{Point{X:2, Y:2}:1, Point{X:2, Y:1}:1}

		v := Vents{}
		ventLines := make([]VentLine,0)
		testReadings :=  []string{"2,2 -> 2,1"}
		for _,s := range testReadings {
			ventLine, err := createVentLine(s)
			assert.Nil(t, err)
			ventLines = append(ventLines, *ventLine)
		}
		fmt.Printf("ventLines is %v\n", ventLines)
		hotSpots := v.FindHotSpots(ventLines)
		fmt.Println(hotSpots)
		assert.Equal(t, expected, hotSpots)
	})
	t.Run("Test Vertical Line intersect", func(t *testing.T) {

		v := Vents{}
		ventLines := make([]VentLine, 0)
		testReadings := []string{"2,2 -> 2,1", "2,1 -> 7,1"}
		for _, s := range testReadings {
			ventLine, err := createVentLine(s)
			assert.Nil(t, err)
			ventLines = append(ventLines, *ventLine)
		}
		fmt.Printf("ventLines is %v\n", ventLines)
		hotSpots := v.FindHotSpots(ventLines)
		fmt.Println(hotSpots)
		result := v.CalculateAnswer(hotSpots)
		assert.Equal(t, 1, result)
	})

	t.Run("Day One", func(t *testing.T) {

		v := Vents{}
		ventLines, err := v.CreateVentLinesFromInput("/Users/jarrett/src/dev/golang/aoc2021/testData/day5/day5test.txt", true)
		assert.Nil(t, err)
		hotSpots := v.FindHotSpots(ventLines)
		fmt.Println(hotSpots)
		result := v.CalculateAnswer(hotSpots)
		assert.Equal(t, 5, result)
	})
}

func TestDayTwoDiagnonals (t *testing.T) {
	t.Run("Diagonals are created", func(t *testing.T) {
		expected := VentLine{
			From: Point{X: 6, Y: 4},
			To:   Point{X: 2, Y: 0},
		}
		s := "6,4 -> 2,0"
		result, err := createVentLine(s)
		assert.Nil(t, err)
		assert.Equal(t, &expected, result)
	})

	t.Run("Test Vertical Line With New Algorithm", func(t *testing.T) {
		expected := map[Point]int{Point{X: 2, Y: 2}: 1, Point{X: 2, Y: 1}: 1}

		v := Vents{}
		ventLines := make([]VentLine, 0)
		testReadings := []string{"2,2 -> 2,1"}
		for _, s := range testReadings {
			ventLine, err := createVentLine(s)
			assert.Nil(t, err)
			ventLines = append(ventLines, *ventLine)
		}
		fmt.Printf("ventLines is %v\n", ventLines)
		hotSpots := v.FindHotSpotsV2(ventLines)
		fmt.Println(hotSpots)
		assert.Equal(t, expected, hotSpots)
	})

	t.Run("Diagonal mark hot spots", func(t *testing.T) {
		expected := map[Point]int{Point{X: 1, Y: 1}: 1, Point{X: 2, Y: 2}: 1, Point{X: 3, Y: 3}: 1}
		v := Vents{}
		ventLines := make([]VentLine, 0)
		testReadings := []string{"1,1 -> 3,3"}
		for _, s := range testReadings {
			ventLine, err := createVentLine(s)
			assert.Nil(t, err)
			ventLines = append(ventLines, *ventLine)
		}
		hotSpots := v.FindHotSpotsV2(ventLines)
		fmt.Printf("hotSpots are %v\n", hotSpots)
		assert.Equal(t, expected, hotSpots)
	})
	t.Run("Diagonal from test data mark hot spots", func(t *testing.T) {
		expected := map[Point]int{Point{X: 6, Y: 4}: 1, Point{X: 5, Y: 3}: 1, Point{X: 4, Y: 2}: 1, Point{X: 3, Y: 1}: 1, Point{X: 2, Y: 0}: 1}
		v := Vents{}
		ventLines := make([]VentLine, 0)
		testReadings := []string{"6,4 -> 2,0"}
		for _, s := range testReadings {
			ventLine, err := createVentLine(s)
			assert.Nil(t, err)
			ventLines = append(ventLines, *ventLine)
		}
		hotSpots := v.FindHotSpotsV2(ventLines)
		fmt.Printf("hotSpots are %v\n", hotSpots)
		assert.Equal(t, expected, hotSpots)
	})

	t.Run("Diagonal from corner to corner", func(t *testing.T) {
		expected := map[Point]int{Point{X: 8, Y: 0}: 1, Point{X: 7, Y: 1}: 1, Point{X: 6, Y: 2}: 1, Point{X: 5, Y: 3}: 1, Point{X: 4, Y: 4}: 1, Point{X: 3, Y: 5}: 1, Point{X: 2, Y: 6}: 1, Point{X: 1, Y: 7}: 1, Point{X: 0, Y: 8}: 1}
		v := Vents{}
		ventLines := make([]VentLine, 0)
		testReadings := []string{"8,0 -> 0,8"}
		for _, s := range testReadings {
			ventLine, err := createVentLine(s)
			assert.Nil(t, err)
			ventLines = append(ventLines, *ventLine)
		}
		hotSpots := v.FindHotSpotsV2(ventLines)
		fmt.Printf("hotSpots are %v\n", hotSpots)
		assert.Equal(t, expected, hotSpots)
	})

	t.Run("Day One with new algo", func(t *testing.T) {

		v := Vents{}
		ventLines, err := v.CreateVentLinesFromInput("/Users/jarrett/src/dev/golang/aoc2021/testData/day5/day5test.txt", true)
		assert.Nil(t, err)
		hotSpots := v.FindHotSpotsV2(ventLines)
		fmt.Println(hotSpots)
		result := v.CalculateAnswer(hotSpots)
		assert.Equal(t, 5, result)
	})
	t.Run("Day Two", func(t *testing.T) {

		v := Vents{}
		ventLines, err := v.CreateVentLinesFromInput("/Users/jarrett/src/dev/golang/aoc2021/testData/day5/day5test.txt", false)
		assert.Nil(t, err)
		hotSpots := v.FindHotSpotsV2(ventLines)
		fmt.Println(hotSpots)
		result := v.CalculateAnswer(hotSpots)
		assert.Equal(t, 12, result)
	})
}
