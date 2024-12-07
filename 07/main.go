package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    operations, err := readInput("input.txt")
    if err != nil {
        panic(err)
    }

    fmt.Printf("operations: %v\n", operations)
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

type operator int

const (
	ADD operator = iota
	SUBTRACT
	MULTIPLY
	DIVIDE
)
