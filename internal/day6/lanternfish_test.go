package day6

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLanternfish(t *testing.T) {
	t.Run("After 3 days", func(t *testing.T) {
		lanternfish := Lanternfish{School: []uint64{3, 4, 3, 1, 2}}

		fmt.Printf("After 0 days: %v\n", lanternfish.School)
		numberOfDays := 3
		lanternfish.SpawnDay1(numberOfDays)

		fmt.Printf("After 3 days: %v\n", lanternfish.School)
		assert.Equal(t, 7, len(lanternfish.School))
	})
	t.Run("After 18 days", func(t *testing.T) {
		expected := []uint64{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 0, 1, 1, 1, 2, 2, 3, 3, 4, 6, 7, 8, 8, 8, 8}
		lanternfish := Lanternfish{School: []uint64{3, 4, 3, 1, 2}}
		lanternfish.SpawnDay1(18)
		fmt.Printf("After 18 days: %v\n", lanternfish.School)
		assert.Equal(t, 26, len(lanternfish.School))
		assert.Equal(t, expected, lanternfish.School)
	})
	t.Run("After 80 days", func(t *testing.T) {
		lanternfish := Lanternfish{School: []uint64{3, 4, 3, 1, 2}}
		lanternfish.SpawnDay1(80)
		assert.Equal(t, 5934, len(lanternfish.School))
	})
	t.Run("After 256 days", func(t *testing.T) {
		lanternfish := Lanternfish{School: []uint64{3, 4, 3, 1, 2}}
		lanternfish.SpawnDay1(256)
		assert.Equal(t, 26984457539, len(lanternfish.School))
	})
}

func TestFastSpawn(t *testing.T) {
	t.Run("After 3 days", func(t *testing.T) {
		l := Lanternfish{}

		initialSchool := []int{3, 4, 3, 1, 2}
		l.SetGradSchoolFromArray(initialSchool)

		fmt.Printf("Day 1 map is %v\n", l.GradSchool)
		l.SpawnDay2(3)
		assert.Equal(t, uint64(7), l.CountFish())
	})
	t.Run("After 18 days", func(t *testing.T) {
		l := Lanternfish{}

		initialSchool := []int{3, 4, 3, 1, 2}
		l.SetGradSchoolFromArray(initialSchool)

		fmt.Printf("Day 1 map is %v\n", l.GradSchool)
		l.SpawnDay2(18)
		assert.Equal(t, uint64(26), l.CountFish())
	})
	t.Run("After 80 days", func(t *testing.T) {
		l := Lanternfish{}

		initialSchool := []int{3, 4, 3, 1, 2}
		l.SetGradSchoolFromArray(initialSchool)

		fmt.Printf("Day 1 map is %v\n", l.GradSchool)
		l.SpawnDay2(80)
		assert.Equal(t, uint64(5934), l.CountFish())
	})
	t.Run("After 256 days", func(t *testing.T) {
		l := Lanternfish{}

		initialSchool := []int{3, 4, 3, 1, 2}
		l.SetGradSchoolFromArray(initialSchool)

		fmt.Printf("Day 1 map is %v\n", l.GradSchool)
		l.SpawnDay2(256)
		assert.Equal(t, uint64(26984457539), l.CountFish())
	})

	t.Run("Getting from file works", func(t *testing.T) {
		l := Lanternfish{}
		initialSchool := []int{3, 4, 3, 1, 2}
		result := l.ConvertInputToArray("/Users/jarrett/src/dev/golang/aoc2021/testData/day6/day6test.txt")
		assert.Equal(t, initialSchool, result)

	})

	t.Run("Getting from full works", func(t *testing.T) {
		l := Lanternfish{}
		result := l.ConvertInputToArray("/Users/jarrett/src/dev/golang/aoc2021/testData/day6/day6.txt")
		assert.Equal(t, 300, len(result))
		l.SetGradSchoolFromArray(result)
		l.SpawnDay2(80)
		assert.Equal(t, uint64(380758), l.CountFish())

	})

	t.Run("256 day answer on all input", func(t *testing.T) {
		l := Lanternfish{}
		result := l.ConvertInputToArray("/Users/jarrett/src/dev/golang/aoc2021/testData/day6/day6.txt")
		assert.Equal(t, 300, len(result))
		l.SetGradSchoolFromArray(result)
		l.SpawnDay2(256)
		fmt.Printf("256 day fish: %d\n", l.CountFish())

	})
}
