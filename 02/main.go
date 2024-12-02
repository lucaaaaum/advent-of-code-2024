package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    input, err := readInput()
    if err != nil {
        panic(err)
    }

    for _, v := range input {
        fmt.Printf("v: %v\n", v)
    }
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

func countSafeReports()  {
    
}
