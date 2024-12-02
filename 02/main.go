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

	safeReportsWithToleranceAmount := countSafeReportsWithTolerance(reports)
	fmt.Printf("safeReportsWithToleranceAmount: %v\n", safeReportsWithToleranceAmount)
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
	safeReports, _ := groupReports(reports)
	return len(safeReports)
}

func groupReports(reports [][]int) ([][]int, [][]int) {
	safeReports := [][]int{}
	unsafeReports := [][]int{}

	for _, report := range reports {
		if isSafe(report) {
			safeReports = append(safeReports, report)
		} else {
			unsafeReports = append(unsafeReports, report)
		}
	}

	return safeReports, unsafeReports
}

func isSafe(report []int) bool {
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

	return safe
}

func countSafeReportsWithTolerance(reports [][]int) int {
	safeReports, unsafeReports := groupReports(reports)

	for _, report := range unsafeReports {
		for itemIndex := range report {
			reportWithoutItem := make([]int, 0)
			reportWithoutItem = append(reportWithoutItem, report[:itemIndex]...)
			reportWithoutItem = append(reportWithoutItem, report[itemIndex+1:]...)
			if isSafe(reportWithoutItem) {
				safeReports = append(safeReports, report)
				break
			}
		}
	}

	return len(safeReports)
}

type Direction int

const (
	Undefined Direction = iota
	Increasing
	Decreasing
)
