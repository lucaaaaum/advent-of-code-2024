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

		if line == "" {
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

type rule struct {
}
