package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules, updates, err := readInput(os.Args[1])
	if err != nil {
		panic(err)
	}

	result := sumMiddlePageNumberForOrderedUpdates(updates, rules)
	fmt.Printf("result: %v\n", result)

	result2 := sumMiddlePageNumberForUnorderedUpdates(updates, rules)
	fmt.Printf("result2: %v\n", result2)
}

func readInput(path string) (map[int][]int, map[int][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := make(map[int][]int)
	updates := make(map[int][]int, 0)

	rulesAreDone := false
	lineCounter := -1
	for scanner.Scan() {
		line := scanner.Text()
		lineCounter++

		if line == "" || line == "\n" {
			rulesAreDone = true
			continue
		}

		if rulesAreDone {
			pagesAsStrings := strings.Split(line, ",")
			for _, v := range pagesAsStrings {
				page, err := strconv.Atoi(v)
				if err != nil {
					return nil, nil, err
				}

				updates[lineCounter] = append(updates[lineCounter], page)
			}
		} else {
			rule := strings.Split(line, "|")

			page, err := strconv.Atoi(rule[0])
			mustPreceedPage, err := strconv.Atoi(rule[1])
			if err != nil {
				return nil, nil, err
			}

			rules[page] = append(rules[page], mustPreceedPage)
		}
	}

	return rules, updates, nil
}

func sumMiddlePageNumberForOrderedUpdates(updates map[int][]int, rules map[int][]int) int {
	result := 0
	for _, update := range updates {
		if isOrdered(update, rules) {
			result += update[len(update)/2]
		}
	}
	return result
}

func isOrdered(update []int, rules map[int][]int) bool {
	for i, page := range update {
		otherPages := make([]int, 0)
		otherPages = append(otherPages, update[:i]...)

		rulesForPage := rules[page]
		for _, otherPage := range otherPages {
			for _, ruleForPage := range rulesForPage {
				if ruleForPage == otherPage {
					return false
				}
			}
		}
	}
	return true
}

func sumMiddlePageNumberForUnorderedUpdates(updates map[int][]int, rules map[int][]int) int {
	result := 0
	for _, update := range updates {
		if !isOrdered(update, rules) {
			update := orderUpdate(update, rules)
			result += update[len(update)/2]
		}
	}
	return result
}

func orderUpdate(update []int, rules map[int][]int) []int {
}
