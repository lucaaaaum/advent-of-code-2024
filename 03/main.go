package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, err := readInput(os.Args[1])
	if err != nil {
		panic(err)
	}

	result1, err := getMultiplicationResult(input)
	fmt.Printf("result: %v\n", result1)

	result2, err := getMultiplicationResultWithEnablingAndDisabling(input)
	fmt.Printf("result2: %v\n", result2)
}

func readInput(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)

	input := ""

	for scanner.Scan() {
		input += scanner.Text()
	}

	return input, nil
}

func getMultiplicationResult(input string) (int, error) {
	r, err := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	if err != nil {
		return -1, err
	}

	matches := r.FindAllStringSubmatch(input, -1)

	result := 0
	for _, match := range matches {
		first, err := strconv.Atoi(match[1])
		if err != nil {
			return -1, err
		}
		second, err := strconv.Atoi(match[2])
		if err != nil {
			return -1, err
		}
		result += first * second
	}

	return result, nil
}

func getMultiplicationResultWithEnablingAndDisabling(input string) (int, error) {
	doRegex, err := regexp.Compile("do\\(\\)")
	dontRegex, err := regexp.Compile("don't\\(\\)")
	mulRegex, err := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	if err != nil {
		return -1, err
	}

	operations := make([]operation, 0)

	enables := doRegex.FindAllStringIndex(input, -1)
	for _, v := range enables {
		operations = append(operations, operation{index: v[0], operationType: Do})
	}
	disables := dontRegex.FindAllStringIndex(input, -1)
	for _, v := range disables {
		operations = append(operations, operation{index: v[0], operationType: Dont})
	}
	muls := mulRegex.FindAllStringIndex(input, -1)
	for _, v := range muls {
		mul := mulRegex.FindStringSubmatch(input[v[0]:v[1]])
		first, err := strconv.Atoi(mul[1])
		if err != nil {
			return -1, err
		}
		second, err := strconv.Atoi(mul[2])
		if err != nil {
			return -1, err
		}
		operations = append(operations, operation{index: v[0], operationType: Mul, first: first, second: second})
	}

	operations = quicksort(operations)

	enabled := true
	result := 0
	for _, operation := range operations {
		switch operation.operationType {
		case Do:
			enabled = true
			break
		case Dont:
			enabled = false
		case Mul:
			if enabled {
				result += operation.first * operation.second
			}
		}
	}

	return result, nil
}

type operationType int

const (
	Do operationType = iota
	Dont
	Mul
)

type operation struct {
	index         int
	operationType operationType
	first         int
	second        int
}

func quicksort(operations []operation) []operation {
	if len(operations) <= 1 {
		return operations
	}

	pivot := operations[len(operations)/2]

	left := make([]operation, 0)
	right := make([]operation, 0)

	for _, v := range operations {
		if v.index < pivot.index {
			left = append(left, v)
		} else if v.index > pivot.index {
			right = append(right, v)
		}
	}

	result := make([]operation, 0)
	result = append(result, quicksort(left)...)
	result = append(result, pivot)
	result = append(result, quicksort(right)...)

	return result
}
