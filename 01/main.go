package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	list1, list2, err := readInput()
	if err != nil {
		panic(err)
	}

	quick1 := quicksort(list1)
	quick2 := quicksort(list2)

	totalDistance := calculateTotalDistance(quick1, quick2)
	fmt.Printf("totalDistance: %v\n", totalDistance)

	similarity := calculateSimilarity(quick1, quick2)
	fmt.Printf("similarity: %v\n", similarity)
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

	left = quicksort(left)
	right = quicksort(right)

	output := append(left, pivot)
	output = append(output, right...)

	return output
}

func calculateTotalDistance(list1 []int, list2 []int) int {
	totalDistance := 0

	for i := 0; i < len(list1); i++ {
		if list1[i] > list2[i] {
			totalDistance += list1[i] - list2[i]
		} else {
			totalDistance += list2[i] - list1[i]
		}
	}

	return totalDistance
}

func calculateSimilarity(list1 []int, list2 []int) int {
	similarities := make([]int, 0)

	lastIndexOfList2 := 0
	for _, v := range list1 {
		reocurrences := 0
		for i := lastIndexOfList2; i < len(list2); i++ {
			if v == list2[i] {
				reocurrences++
				lastIndexOfList2 = i
			} else if v < list2[i] {
				break
			}
		}
		similarities = append(similarities, v*reocurrences)
	}

	totalSimilarities := 0
	for _, v := range similarities {
		totalSimilarities += v
	}
	return totalSimilarities
}
