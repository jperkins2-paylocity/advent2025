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

		rotation := strings.Fields(line)

		for _, r := range rotation {
			direction := r[0]
			value, err := strconv.Atoi(r[1:])
			if err != nil {
				panic(fmt.Sprintf("Error converting value: %v\n", err))
			}

			fullRotations += value / totalDialPositions
			clicks := value % totalDialPositions

			switch direction {
			case 'R':
				currentDialPosition += clicks
				if currentDialPosition > maxDialValue {
					currentDialPosition %= totalDialPositions
					if currentDialPosition != 0 {
						fullRotations++
					}
				}
			case 'L':
				currentDialPosition -= clicks
				if currentDialPosition < minDialValue {
					currentDialPosition = (currentDialPosition + totalDialPositions)
					if currentDialPosition != 0 {
						fullRotations++
					}
				}
			default:
				panic("invalid direction")
			}

			if currentDialPosition == 0 {
				zeroCount++
			}

			currentDialPosition = (currentDialPosition + totalDialPositions) % totalDialPositions

		}
	}

	fmt.Printf("Part1 Zero Count: %d\n", zeroCount)
	fmt.Printf("Part2 Total Zero Count: %d\n", (zeroCount + fullRotations))
}
