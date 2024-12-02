package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	list1, list2, err := readInput()
	if err != nil {
		panic(err)
	}
}

func readInput() ([]int, []int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return []int{}, []int{}, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	ids1, ids2 := []int{}, []int{}
	for scanner.Scan() {
		line := scanner.Text()
		ids := strings.Split(line, "   ")
		id1, err := strconv.Atoi(ids[0])
		if err != nil {
			return []int{}, []int{}, err
		}
		id2, err := strconv.Atoi(ids[1])
		if err != nil {
			return []int{}, []int{}, err
		}
		ids1 = append(ids1, id1)
		ids2 = append(ids2, id2)
	}

	return ids1, ids2, nil
}

func quicksort(input []int) []int {
	if len(input) <= 1 {
		return input
	}

	pivot := input[0]

	left := []int{}
	right := []int{}

	for i := 1; i < len(input); i++ {
		if input[i] < pivot {
			left = append(left, input[i])
		} else {
			right = append(right, input[i])
		}
	}

	output := append(left, pivot)
	output = append(output, right...)

	return output
}
