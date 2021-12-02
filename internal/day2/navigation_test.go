package day2

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_DayTwo(t *testing.T) {

    t.Run("Empty array", func(t *testing.T) {
        testInput := []Navigation{}
        n := NavigationSystem{}
        result, err := n.CalculatePosition(testInput)
        assert.Nil(t, err)
        assert.Equal(t, Position{}, result)
    })
    t.Run("Bad direction throws an error", func(t *testing.T) {
        testInput := []Navigation{{
            Direction: "foo",
            Position:  1,
        }}
        n := NavigationSystem{}
        _, err := n.CalculatePosition(testInput)
        assert.NotNil(t, err)
    })
    t.Run("Happy Path", func(t *testing.T) {
        nav1 := Navigation{
            Direction: "forward",
            Position:  5,
        }
        nav2 := Navigation{
            Direction: "down",
            Position:  5,
        }
        testInput := []Navigation{nav1, nav2}
        n := NavigationSystem{}
        result, err := n.CalculatePosition(testInput)
        assert.Nil(t, err)
        assert.Equal(t, Position{Horizontal: 5, Depth: 5}, result)
    })
    t.Run("Day 2 test", func(t *testing.T) {
        testInput := []Navigation{Navigation{Direction: "forward", Position: 5}, Navigation{Direction: "down", Position: 5}, Navigation{Direction: "forward", Position: 8}, Navigation{Direction: "up", Position: 3}, Navigation{Direction: "down", Position: 8}, Navigation{Direction: "forward", Position: 2}}
        n := NavigationSystem{}
        result, err := n.CalculatePosition(testInput)
        assert.Nil(t, err)
        assert.Equal(t, Position{Horizontal: 15, Depth: 10}, result)
    })
}
