package application

import (
	"bufio"
	"fmt"
	"github.com/jrjarrett/aoc2021/internal/day1"
	"log"
	"os"
	"strconv"
)

type App struct {
	Sonar day1.Sonar
}

func New() *App {
	application := App{Sonar: day1.Sonar{}}
	return &application
}

func (a *App) Run() {
	a.Day1_1()
	a.Day1_2()
}

func (a *App) Day1_1() {
	sonarSweep, err := getAOCSonarInput("testData/day1/challenge1.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := a.Sonar.DepthIncrease(sonarSweep)
	fmt.Printf("Day 1.1 result is %#v\n", result)
}

func (a *App) Day1_2() {
	sonarSweep, err := getAOCSonarInput("testData/day1/challenge1.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := a.Sonar.DepthIncrease(a.Sonar.CreateSlidingWindows(sonarSweep))
	fmt.Printf("Day 1.2 result is %#v\n", result)
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
