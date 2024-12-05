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

	occurrences := grid.countXmasOccurences()
	fmt.Printf("occurrences: %v\n", occurrences)
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
	items  [][]rune
	values [][]string
}

func newGrid() grid {
	return grid{items: make([][]rune, 0)}
}

func (g grid) countXmasOccurences() int {
	countXmas := 0
	for y, line := range g.items {
		for x, char := range line {
			if char == 'X' {
				fmt.Printf("Found X at (%v, %v)\n", x, y)
				countXmas += g.countXmas(x, y)
			}
		}
	}
	return countXmas
}

func (g grid) countXmas(x, y int) int {
	count := 0
	directions := []direction{Top, TopRight, Right, BottomRight, Bottom, BottomLeft, Left, TopLeft}
	for _, dir := range directions {
		if g.findSequence("MAS", x, y, dir) {
			fmt.Printf("Found XMAS at (%v, %v) going to %v\n", x, y, dir)
			count++
		}
	}
	return count
}

func (g grid) findSequence(sequence string, x, y int, dir direction) bool {
	x, y = applyDirection(x, y, dir)

	if x < 0 || y < 0 || y >= len(g.items) || x >= len(g.items[y]) {
		return false
	}

	current := rune(sequence[0])

	if g.items[y][x] != current {
		return false
	}

	if len(sequence) == 1 {
		return true
	}

	return g.findSequence(sequence[1:], x, y, dir)
}

func applyDirection(x, y int, dir direction) (int, int) {
	switch dir {
	case Top:
		y--
		break
	case TopRight:
		y--
		x++
		break
	case Right:
		x++
		break
    case BottomRight:
        y++
        x++
	case Bottom:
		y++
		break
	case BottomLeft:
		y++
		x--
		break
	case Left:
		x--
		break
	case TopLeft:
		y--
		x--
		break
	}

	return x, y
}

type direction int

const (
	Top direction = iota
	TopRight
	Right
	BottomRight
	Bottom
	BottomLeft
	Left
	TopLeft
)
