package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Lanternfish struct {
	School     []uint64
	GradSchool map[int]uint64
}

func (l *Lanternfish) SetIntialSchoolFromFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	vals := strings.Split(scanner.Text(), ",")
	for i := 0; i < len(vals); i++ {
		fish, err := strconv.Atoi(vals[i])
		if err != nil {
			panic("Bad input from file")
		}
		l.School = append(l.School, uint64(fish))
	}
	fmt.Printf("Number of inputs is %d\n", len(l.School))
}

func (l *Lanternfish) ConvertInputToArray(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	result := make([]int, 0)
	vals := strings.Split(scanner.Text(), ",")
	for i := 0; i < len(vals); i++ {
		fish, err := strconv.Atoi(vals[i])
		if err != nil {
			panic("Bad input from file")
		}
		result = append(result, fish)
	}
	return result
}
func (l *Lanternfish) SpawnDay1(numberOfDays int) {
	var numberSpawned int
	for i := 1; i < numberOfDays; i++ {
		if i%10 == 0 {
			fmt.Printf("Day %d, fish %d\n", i, len(l.School))
		}

		for fish := 0; fish < len(l.School); fish++ {
			if l.School[fish] == 0 {
				l.School[fish] = 6
				numberSpawned++
			} else {
				l.School[fish] = l.School[fish] - 1
			}
		}
		if numberSpawned > 0 {
			// fmt.Printf("On day %d, spawned %d fish\n", i, numberSpawned)
			for n := 0; n < numberSpawned; n++ {
				l.School = append(l.School, 8)
			}
		}
		numberSpawned = 0
	}
}
func (l *Lanternfish) SetGradSchoolFromArray(initialSchool []int) {

	l.GradSchool = make(map[int]uint64)
	// Seed the initial map
	for i := 0; i <= 8; i++ {
		l.GradSchool[i] = 0
	}
	for i := 0; i < len(initialSchool); i++ {
		l.GradSchool[initialSchool[i]] = l.GradSchool[initialSchool[i]] + 1
	}
}
func (l *Lanternfish) SpawnDay2(numberOfDays int) {
	for days := 1; days <= numberOfDays; days++ {
		// Process a day
		numberToSpawn := l.GradSchool[0]
		for i := 0; i < len(l.GradSchool); i++ {
			l.GradSchool[i] = l.GradSchool[i+1]
		}
		l.GradSchool[8] = l.GradSchool[8] + numberToSpawn
		l.GradSchool[6] = l.GradSchool[6] + numberToSpawn
	}
}

func (l *Lanternfish) CountFish() uint64 {
	var totalFish uint64
	for i := 0; i <= 8; i++ {
		totalFish = totalFish + l.GradSchool[i]
	}
	return totalFish
}
