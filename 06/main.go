package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	g, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range g.fields {
		for _, field := range line {
			switch field {
			case empty:
				fmt.Print(".")
				break
			case obstacle:
				fmt.Print("#")
				break
			case guard:
				fmt.Print("^")
				break
			}
		}
		fmt.Println()
	}
}

type grid struct {
	fields [][]field
}

type field int

const (
	empty field = iota
	obstacle
	guard
)

func readInput(path string) (grid, error) {
	file, err := os.Open(path)
	if err != nil {
		return grid{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	g := grid{fields: make([][]field, 0)}

	currentLine := -1
	for scanner.Scan() {
        currentLine++
		line := scanner.Text()
        fmt.Printf("line: %v\n", line)
		g.fields = append(g.fields, make([]field, len(line)))
		for currentField, field := range line {
			switch field {
			case '.':
				g.fields[currentLine][currentField] = empty
				break
			case '#':
				g.fields[currentLine][currentField] = obstacle
				break
			case '^':
				g.fields[currentLine][currentField] = guard
				break
			default:
				return grid{}, fmt.Errorf("Unknown field type: %v", field)
			}
		}
	}

	return g, nil
}
