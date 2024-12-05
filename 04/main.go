package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid, err := readInput(os.Args[1])
	if err != nil {
		panic(err)
	}

    for _, v := range grid.items {
        for _, v := range v {
            fmt.Print(string(v))
        }
        fmt.Println()
    }
}

func readInput(path string) (grid, error) {
	file, err := os.Open(path)
	if err != nil {
		return grid{}, err
	}

	scanner := bufio.NewScanner(file)
	grid := grid{}
	currentLine := 0

	for scanner.Scan() {
		line := scanner.Text()
		grid.items = append(grid.items, make([]rune, 0))
		for _, v := range line {
			grid.items[currentLine] = append(grid.items[currentLine], v)
		}
		currentLine++
	}

	return grid, nil
}

type grid struct {
	items [][]rune
}

func newGrid() grid {
	return grid{items: make([][]rune, 0)}
}
