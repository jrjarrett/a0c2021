package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DayOne(t *testing.T) {

	t.Run("Empty array", func(t *testing.T) {
		testInput := []int{}
		s := Sonar{}
		result := s.DepthIncrease(testInput)

		assert.Equal(t, 0, result)
	})
	t.Run("One Increase", func(t *testing.T) {
		testInput := []int{199, 200}
		s := Sonar{}
		result := s.DepthIncrease(testInput)

		assert.Equal(t, 1, result)
	})

	t.Run("AOC example", func(t *testing.T) {
		testInput := []int{199,
			200,
			208,
			210,
			200,
			207,
			240,
			269,
			260,
			263}
		s := Sonar{}
		result := s.DepthIncrease(testInput)

		assert.Equal(t, 7, result)
	})
	t.Run("test SlidingWindows with fewer than 3 readings", func(t *testing.T) {
		readings := []int{199, 200}
		s := Sonar{}
		result := s.CreateSlidingWindows(readings)
		assert.Equal(t, 0, len(result))
		//assert.Equal(t, 607, result[0])
	})
	t.Run("test SlidingWindows", func(t *testing.T) {
		readings := []int{199, 200, 208}
		s := Sonar{}
		result := s.CreateSlidingWindows(readings)
		assert.Equal(t, 1, len(result))
		assert.Equal(t, 607, result[0])
	})
	t.Run("test SlidingWindows with AOC example", func(t *testing.T) {
		readings := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
		s := Sonar{}
		result := s.CreateSlidingWindows(readings)
		assert.Equal(t, 8, len(result))
		assert.Equal(t, 607, result[0])
		assert.Equal(t, 647, result[4])
		assert.Equal(t, 792, result[7])
	})
}
