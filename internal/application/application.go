package application

import (
	"bufio"
	"errors"
	"fmt"
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
}

func New() *App {
	application := App{Sonar: day1.Sonar{}}
	return &application
}

func (a *App) Run() {
	//a.Day1_1()
	//a.Day1_2()
	//a.Day2_1()
	//a.Day2_2()
	a.Day3()

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

func (a *App) Day3() {
	readings, err := a.Scanner.GetReactorInput("/Users/jarrett/src/aoc2021/testData/day3/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	transform := a.Scanner.RotateArrayCW(readings)
	gamma, epsilon := a.Scanner.CalculateGammaEpsilonValue(transform)
	fmt.Printf("Day 2.1 Gamma is %d, epsilon is %d, energy use is %d\n", gamma, epsilon, gamma*epsilon)
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
