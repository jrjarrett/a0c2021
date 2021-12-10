package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bingo struct {
	Draws  []int
	Boards []Board
}

type BingoNumber struct {
	row    int
	col    int
	marked bool
}

// Numbers on the board are stored in a map of called numbers to their position.
// Consider:
//          22 13 17 11  0
//          8  2 23  4 24
//          21  9 14 16  7
//          6 10  3 18  5
//          1 12 20 15 19
// Then board[22] = BingoNumber{0,0}
//      board[2] = BingoNumber{1,1}
//      board[19] = BingoNumber{4,4}
// Call a number, then we add the BingoNumber.row to rowBingo, BingoNumber.col to columnBingo. If rowBingo == rows, BINGO!
// if columnBingo == cols, BINGO!

type Board struct {
	ID          int
	rows        int
	cols        int
	marked      map[int]BingoNumber
	rowBingo    []int
	columnBingo []int
}

func (b *Bingo) BuildGameFromInput(fileName string) Bingo {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	draws, err := convertToInts(strings.Split(scanner.Text(), ","))

	if err != nil {
		panic("Bad data in draws line of input!")
	}

	// skip blank line between number and boards
	scanner.Scan()
	scanner.Text()

	boards := make([]Board, 0)
	linesForBoard := make([]string, 0)
	boardId := 1
	for scanner.Scan() {
		if len(linesForBoard) == 5 {
			boards = append(boards, NewBoard(linesForBoard, boardId))
			boardId++
			linesForBoard = make([]string, 0)
			scanner.Text()
		} else {
			linesForBoard = append(linesForBoard, scanner.Text())
		}
	}
	boards = append(boards, NewBoard(linesForBoard, boardId))

	return Bingo{Draws: draws, Boards: boards}
}

func NewBoard(linesInBoard []string, boardId int) Board {
	board := Board{
		ID:          boardId,
		rows:        5,
		cols:        5,
		marked:      make(map[int]BingoNumber),
		rowBingo:    make([]int, 5),
		columnBingo: make([]int, 5),
	}
	// We now have 5 lines of the new board as strings.
	for rowCount, line := range linesInBoard {
		line = strings.TrimSpace(strings.Replace(line, "  ", " ", -1))
		linex := strings.Split(line, " ")
		lineValues, err := convertToInts(linex)
		// lineValues are the values in row "rowCount" of the board.
		if err != nil {
			panic("Failed to build new board")
		}
		for colCount, value := range lineValues {
			bn := BingoNumber{rowCount, colCount, false}
			board.marked[value] = bn
		}

	}
	return board
}

func ApplyDrawsToBoards(b Bingo) (*Board, int) {
	for i, draw := range b.Draws {
		fmt.Printf("Draw # %d - draws %d\n", i+1, draw)
		for i := range b.Boards {
			if _, ok := b.Boards[i].marked[draw]; ok {
				b.Boards[i].rowBingo[b.Boards[i].marked[draw].row]++
				b.Boards[i].columnBingo[b.Boards[i].marked[draw].col]++
				bn := b.Boards[i].marked[draw]
				bn.marked = true
				b.Boards[i].marked[draw] = bn
				if b.Boards[i].rowBingo[b.Boards[i].marked[draw].row] == b.Boards[i].rows || b.Boards[i].columnBingo[b.Boards[i].marked[draw].col] == b.Boards[i].cols {
					// BINGO!
					fmt.Printf("BINGO! on board %d for draw %d\n", b.Boards[i].ID, draw)
					return &b.Boards[i], draw
				}
			}
		}
		fmt.Printf("Gone thru all the boards for draw %d, no winner\n\n", i+1)
	}
	fmt.Println("Something wrong - no winners?")
	return nil, -1
}
func ApplyDrawsToBoardsV2(b Bingo) (*Board, int) {
	var lastWinningBoard Board
	var lastDraw int
	for drawNumber, draw := range b.Draws {
		bingoedBoards := make([]int, 0)
		bingoedBoardIDs := make([]int, 0)
		fmt.Printf("Draw # %d - draws %d\n", drawNumber+1, draw)
		for i := 0; i < len(b.Boards); i++ {
			if _, ok := b.Boards[i].marked[draw]; ok {
				b.Boards[i].rowBingo[b.Boards[i].marked[draw].row]++
				b.Boards[i].columnBingo[b.Boards[i].marked[draw].col]++
				bn := b.Boards[i].marked[draw]
				bn.marked = true
				b.Boards[i].marked[draw] = bn
				if b.Boards[i].rowBingo[b.Boards[i].marked[draw].row] == b.Boards[i].rows || b.Boards[i].columnBingo[b.Boards[i].marked[draw].col] == b.Boards[i].cols {
					// BINGO!
					fmt.Printf("BINGO! on board %d for draw %d\n", b.Boards[i].ID, draw)
					lastWinningBoard = b.Boards[i]
					lastDraw = draw
					bingoedBoards = append(bingoedBoards, i)
					bingoedBoardIDs = append(bingoedBoardIDs, b.Boards[i].ID)
				}
			}
		}
		// If we had bingos, remove that board.
		if len(bingoedBoards) > 0 {
			for bbid, bingoedBoard := range bingoedBoards {
				fmt.Printf("Removing board number %d board ID is %d\n", bingoedBoard, bingoedBoardIDs[bbid])
				if bingoedBoard < len(b.Boards) {
					b.Boards = append(b.Boards[:bingoedBoard], b.Boards[bingoedBoard+1:]...)
				} else {
					b.Boards = append(b.Boards[:bingoedBoard])
				}
			}

			// Reset for next draw
			bingoedBoards = make([]int, 0)
		}
	}
	return &lastWinningBoard, lastDraw
}
func CalculateUnmarkedSquares(b Board) int {
	var boardSum int
	for key, element := range b.marked {
		if !element.marked {
			boardSum = boardSum + key

		}
	}
	return boardSum
}

func convertToInts(strs []string) ([]int, error) {
	draws := make([]int, len(strs))
	for counter, val := range strs {
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		draws[counter] = i
	}
	return draws, nil
}
