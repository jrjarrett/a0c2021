package application

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/jrjarrett/aoc2021/internal/day4"
	"github.com/jrjarrett/aoc2021/internal/day5"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jrjarrett/aoc2021/internal/day1"
	"github.com/jrjarrett/aoc2021/internal/day2"
	"github.com/jrjarrett/aoc2021/internal/day3"
)

type App struct {
	Sonar            day1.Sonar
	NavigationSystem day2.NavigationSystem
	Scanner          day3.Scanner
	Bingo            day4.Bingo
	Vents            day5.Vents
}

func New() *App {
	application := App{Sonar: day1.Sonar{}}
	return &application
}

func (a *App) Run() {
	// a.Day1_1()
	// a.Day1_2()
	// a.Day2_1()
	// a.Day2_2()
	// a.Day3_1()
	// a.Day3_2()
	// a.Day4_1()
	// a.Day4_2()
	a.Day5_1()
	a.Day5_2()

}

func (a *App) Day1_1() {

	sonarSweep, err := getAOCSonarInput("/Users/jarrett/src/aoc2021/testData/day1/challenge1.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := a.Sonar.DepthIncrease(sonarSweep)
	fmt.Printf("Day 1.1 result is %#v\n", result)
}

func (a *App) Day1_2() {
	sonarSweep, err := getAOCSonarInput("/Users/jarrett/src/aoc2021/testData/day1/challenge1.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := a.Sonar.DepthIncrease(a.Sonar.CreateSlidingWindows(sonarSweep))
	fmt.Printf("Day 1.2 result is %#v\n", result)
}

func (a *App) Day2_1() {
	position, err := getAOCDirectionInput("/Users/jarrett/src/aoc2021/testData/day2/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	result, err := a.NavigationSystem.CalculatePositionV1(position)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Day 2.1 location is %#v. Answer is %d\n", result, result.Horizontal*result.Depth)
}

func (a *App) Day2_2() {
	position, err := getAOCDirectionInput("/Users/jarrett/src/aoc2021/testData/day2/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	result, err := a.NavigationSystem.CalculatePositionV2(position)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Day 2.2 location is %#v. Answer is %d\n", result, result.Horizontal*result.Depth)
}

func (a *App) Day3_1() {
	readings, err := a.Scanner.GetReactorInput("/Users/jarrett/src/aoc2021/testData/day3/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	transform := a.Scanner.RotateArrayCW(readings)
	gamma, epsilon := a.Scanner.CalculateGammaEpsilonValue(transform)
	fmt.Printf("Day 2.1 Gamma is %d, epsilon is %d, energy use is %d\n", gamma, epsilon, gamma*epsilon)
}

func (a *App) Day3_2() {
	readings, err := a.Scanner.GetReactorInput("/Users/jarrett/src/aoc2021/testData/day3/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	o2Readings := a.Scanner.CalculateO2CO2Values(readings, 0, day3.O2)
	co2Readings := a.Scanner.CalculateO2CO2Values(readings, 0, day3.CO2)
	o2 := a.Scanner.ConvertToDecimal(o2Readings)
	co2 := a.Scanner.ConvertToDecimal(co2Readings)
	fmt.Printf("Day 3.2 O2 level is %d CO2 level is %d Life Support Reading is %d", o2, co2, o2*co2)

}

func (a *App) Day4_1() {
	bingo := a.Bingo.BuildGameFromInput("/Users/jarrett/src/dev/golang/aoc2021/testData/day4/day4.txt")
	winningBoard, winningDraw := day4.ApplyDrawsToBoards(bingo)

	boardSum := day4.CalculateUnmarkedSquares(*winningBoard)
	fmt.Printf("sum is %d\n", boardSum)
	fmt.Printf("Answer is is %d\n", boardSum*winningDraw)

}

func (a *App) Day4_2() {
	bingo := a.Bingo.BuildGameFromInput("/Users/jarrett/src/dev/golang/aoc2021/testData/day4/day4.txt")
	winningBoard, winningDraw := day4.ApplyDrawsToBoardsV2(bingo)

	boardSum := day4.CalculateUnmarkedSquares(*winningBoard)
	fmt.Printf("board id is %d\n", winningBoard.ID)
	fmt.Printf("sum is %d\n", boardSum)
	fmt.Printf("Answer is is %d\n", boardSum*winningDraw)

}

func (a *App) Day5_1() {
	ventLines, err := a.Vents.CreateVentLinesFromInput("/Users/jarrett/src/dev/golang/aoc2021/testData/day5/day5.txt", true)
	if err != nil {
		log.Fatal(err)
	}
	hotSpots := a.Vents.FindHotSpotsV2(ventLines)
	result := a.Vents.CalculateAnswer(hotSpots)
	fmt.Printf("Day 5 part 1 answer is %d\n", result)
}

func (a *App) Day5_2() {
	ventLines, err := a.Vents.CreateVentLinesFromInput("/Users/jarrett/src/dev/golang/aoc2021/testData/day5/day5.txt", false)
	if err != nil {
		log.Fatal(err)
	}
	hotSpots := a.Vents.FindHotSpotsV2(ventLines)
	result := a.Vents.CalculateAnswer(hotSpots)
	fmt.Printf("Day 5 part 2 answer is %d\n", result)
}

func getAOCSonarInput(fileName string) ([]int, error) {
	var sonarInput []int
	file, err := os.Open(fileName)
	if err != nil {
		return sonarInput, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 0)
		if err != nil {
			return sonarInput, err
		}
		sonarInput = append(sonarInput, int(i))
	}
	err = file.Close()
	if err != nil {
		return sonarInput, err
	}
	if scanner.Err() != nil {
		return sonarInput, scanner.Err()
	}
	return sonarInput, nil
}

func getAOCDirectionInput(fileName string) ([]day2.Navigation, error) {
	var line string
	var maneuvers []day2.Navigation
	file, err := os.Open(fileName)
	if err != nil {
		return maneuvers, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		if len(strings.Fields(line)) != 2 {
			return maneuvers, errors.New("Wrong input in file " + line)
		}
		direction := strings.Fields(line)[0]
		position, err := strconv.Atoi(strings.Fields(line)[1])
		if err != nil {
			return maneuvers, err
		}
		nav := day2.Navigation{
			Direction: direction,
			Position:  position,
		}
		maneuvers = append(maneuvers, nav)
	}
	return maneuvers, nil
}
