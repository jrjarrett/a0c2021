package day4

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DayFourPart1(t *testing.T) {
	t.Run("test AOC example builds boards correctly", func(t *testing.T) {
		b := Bingo{}
		b = b.BuildGameFromInput("/Users/jarrett/src/aoc2021/testData/day4/day4test.txt")
		assert.NotNil(t, b)
		assert.Equal(t, 27, len(b.Draws))
		assert.Equal(t, 3, len(b.Boards))
		assert.Equal(t, BingoNumber{row: 2, col: 3}, b.Boards[0].marked[16])
		fmt.Printf("Number of numbers to draw is %#v\n\n number of boards is %#v\n", len(b.Draws), len(b.Boards))

	})
	t.Run("test AOC example plays correctly", func(t *testing.T) {
		b := Bingo{}
		b = b.BuildGameFromInput("/Users/jarrett/src/aoc2021/testData/day4/day4test.txt")
		winningBoard, draw := ApplyDrawsToBoards(b)
		assert.Equal(t, draw, 24)
		assert.Equal(t, BingoNumber{row: 0, col: 0, marked: true}, winningBoard.marked[14])
		assert.Equal(t, BingoNumber{row: 0, col: 1, marked: true}, winningBoard.marked[21])
		assert.Equal(t, BingoNumber{row: 0, col: 2, marked: true}, winningBoard.marked[17])
		assert.Equal(t, BingoNumber{row: 0, col: 3, marked: true}, winningBoard.marked[24])
		assert.Equal(t, BingoNumber{row: 0, col: 4, marked: true}, winningBoard.marked[4])

	})
	t.Run("test AOC example final score", func(t *testing.T) {
		b := Bingo{}
		b = b.BuildGameFromInput("/Users/jarrett/src/aoc2021/testData/day4/day4test.txt")
		winningBoard, winningDraw := ApplyDrawsToBoards(b)
		assert.Equal(t, winningDraw, 24)
		assert.Equal(t, BingoNumber{row: 0, col: 0, marked: true}, winningBoard.marked[14])
		assert.Equal(t, BingoNumber{row: 0, col: 1, marked: true}, winningBoard.marked[21])
		assert.Equal(t, BingoNumber{row: 0, col: 2, marked: true}, winningBoard.marked[17])
		assert.Equal(t, BingoNumber{row: 0, col: 3, marked: true}, winningBoard.marked[24])
		assert.Equal(t, BingoNumber{row: 0, col: 4, marked: true}, winningBoard.marked[4])
		boardSum := CalculateUnmarkedSquares(*winningBoard)
		fmt.Printf("sum is %d\n", boardSum)
		assert.Equal(t, 4512, (boardSum * winningDraw))

	})
	t.Run("test AOC example part 2", func(t *testing.T) {
		b := Bingo{}
		b = b.BuildGameFromInput("/Users/jarrett/src/aoc2021/testData/day4/day4test.txt")
		winningBoard, winningDraw := ApplyDrawsToBoardsV2(b)
		assert.Equal(t, winningDraw, 13)
		assert.Equal(t, BingoNumber{row: 0, col: 2, marked: true}, winningBoard.marked[0])
		assert.Equal(t, BingoNumber{row: 1, col: 2, marked: true}, winningBoard.marked[13])
		assert.Equal(t, BingoNumber{row: 2, col: 2, marked: true}, winningBoard.marked[7])
		assert.Equal(t, BingoNumber{row: 3, col: 2, marked: true}, winningBoard.marked[10])
		assert.Equal(t, BingoNumber{row: 4, col: 2, marked: true}, winningBoard.marked[16])
		boardSum := CalculateUnmarkedSquares(*winningBoard)
		fmt.Printf("sum is %d\n", boardSum)
		assert.Equal(t, 1924, (boardSum * winningDraw))

	})
}
