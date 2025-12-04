package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var totalDialPositions = 100
var maxDialValue = 99
var minDialValue = 0

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("day1/input1.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	zeroCount := 0
	currentDialPosition := 50

	for scanner.Scan() {
		line := scanner.Text()

		rotationList := strings.Fields(line)

		for _, r := range rotationList {
			direction := r[0]
			value, err := strconv.Atoi(r[1:])
			if err != nil {
				panic(fmt.Sprintf("Error converting value: %v\n", err))
			}

			clicks := value % totalDialPositions

			temp := 0

			switch direction {
			case 'R':
				temp = currentDialPosition + clicks
				currentDialPosition = temp
			case 'L':
				temp = currentDialPosition - clicks
				currentDialPosition = temp
			default:
				panic("invalid direction")
			}

			currentDialPosition = (currentDialPosition + totalDialPositions) % totalDialPositions

			if currentDialPosition == 0 {
				zeroCount++
			}
		}
	}

	fmt.Printf("Part1 Password: %d\n", zeroCount)
}

func part2() {
	file, err := os.Open("day1/input1.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	zeroCount := 0
	fullRotations := 0
	currentDialPosition := 50

	for scanner.Scan() {
		line := scanner.Text()

		rotationList := strings.Fields(line)

		for _, r := range rotationList {
			direction := r[0]
			value, err := strconv.Atoi(r[1:])
			if err != nil {
				panic(fmt.Sprintf("Error converting value: %v\n", err))
			}

			fullRotations += value / totalDialPositions
			clicks := value % totalDialPositions

			temp := 0

			switch direction {
			case 'R':
				temp = currentDialPosition + clicks
				if currentDialPosition != 0 && temp > maxDialValue {
					fullRotations++
				}
				currentDialPosition = temp
			case 'L':
				temp = currentDialPosition - clicks
				if currentDialPosition != 0 && temp < minDialValue {
					fullRotations++
				}
				currentDialPosition = temp
			default:
				panic("invalid direction")
			}

			if currentDialPosition == 0 {
				zeroCount++
			}

			currentDialPosition = (currentDialPosition + totalDialPositions) % totalDialPositions
		}
	}

	fmt.Printf("Part2 Password: %d\n", (zeroCount + fullRotations))
}
