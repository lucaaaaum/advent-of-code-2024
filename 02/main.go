package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports, err := readInput()
	if err != nil {
		panic(err)
	}

	safeReportsAmount := countSafeReports(reports)
	fmt.Printf("safeReportsAmount: %v\n", safeReportsAmount)
}

func readInput() ([][]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return [][]int{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")

		itemsAsInt := []int{}
		for _, v := range items {
			itemAsInt, err := strconv.Atoi(v)
			if err != nil {
				return [][]int{}, err
			}
			itemsAsInt = append(itemsAsInt, itemAsInt)
		}
		input = append(input, itemsAsInt)
	}

	return input, nil
}

func countSafeReports(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		safe := true
		direction := Undefined

		for i, current := range report {
			if i != len(report)-1 {
				next := report[i+1]
				currentDirection := Undefined
				difference := 0

				if current < next {
					currentDirection = Increasing
					difference = next - current
				} else {
					currentDirection = Decreasing
					difference = current - next
				}

				if difference < 1 || difference > 3 {
					safe = false
					break
				}

				if direction == Undefined {
					direction = currentDirection
				}

				if currentDirection != direction {
					safe = false
					break
				}
			}
		}

		if safe {
			safeReports++
		}
	}

	return safeReports
}

type Direction int

const (
	Undefined Direction = iota
	Increasing
	Decreasing
)
