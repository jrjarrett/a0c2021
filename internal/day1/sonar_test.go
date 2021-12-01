package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DayOne(t *testing.T) {

	t.Run("One Increase", func(t *testing.T) {
		testInput := []int{199,200}
		s := Sonar{}
		result := s.DepthIncrease(testInput)

		assert.Equal(t, 1, result)
	})
