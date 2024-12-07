package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	operations, err := readInput(os.Args[1])
	if err != nil {
		panic(err)
	}

	count := 0
	for _, v := range operations {
		if v.solve() {
			count += v.target
		}
	}
	fmt.Printf("count: %v\n", count)
}

func readInput(path string) ([]operation, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	operations := make([]operation, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		halfs := strings.Split(line, ":")
		target, err := strconv.Atoi(halfs[0])
		if err != nil {
			return nil, err
		}

		valuesAsString := strings.Split(halfs[1], " ")
		values := make([]int, 0)
		for _, v := range valuesAsString {
			if v == "" {
				continue
			}
			value, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			values = append(values, value)
		}

		operations = append(operations, operation{target: target, values: values})
	}

	return operations, nil
}

type operation struct {
	target int
	values []int
}

func (o operation) solve() bool {
	operators := make([]operator, len(o.values[1:]))
	reset(operators)

	index := 0
	for true {
		o.printCalculus(operators)
		if o.target == o.calculate(operators) {
			return true
		}

		i, err := switchToNextCombination(operators, index)
		index = i

		if err != nil {
			return false
		}
	}

	return false
}

func (o operation) printCalculus(operators []operator) {
	fmt.Printf("%v: %v ", o.target, o.values[0])
	for i, operator := range operators {
		if operator == ADD {
			fmt.Printf("\033[34m+\033[0m %v ", o.values[i+1])
		} else {
			fmt.Printf("\033[33m*\033[0m %v ", o.values[i+1])
		}
	}
	fmt.Printf("= %v", o.calculate(operators))
	fmt.Println()
}

func switchToNextCombination(operators []operator, index int) (int, error) {
	empty := true
	for _, o := range operators {
		if o == MULTIPLY {
            empty = false
			break
		}
	}
	if empty {
		operators[0] = MULTIPLY
		return 0, nil
	}

	if index == len(operators)-1 || operators[index+1] == MULTIPLY {
		if operators[0] == ADD {
			operators[0] = MULTIPLY
			return 0, nil
		}
		return -1, fmt.Errorf("no more combinations")
	}

	operators[index+1] = MULTIPLY
	operators[index] = ADD
	return index + 1, nil
}

func (o operation) calculate(operators []operator) int {
	result := o.values[0]
	for i, operator := range operators {
		if operator == ADD {
			result += o.values[i+1]
		} else {
			result *= o.values[i+1]
		}
	}

	return result
}

func reset(operators []operator) {
	for i := range operators {
		operators[i] = ADD
	}
}

type operator int

const (
	ADD operator = iota
	MULTIPLY
)

func (o operator) invert() operator {
	if o == ADD {
		o = MULTIPLY
	} else {
		o = ADD
	}
	return o
}
