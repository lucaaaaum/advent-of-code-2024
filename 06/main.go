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

	g.run()
	g.printGrid(up)

	countVisited := g.countVisitedFields()
	fmt.Printf("countVisited: %v\n", countVisited)

	loops, err := g.countLoops()
	if err != nil {
		panic(err)
	}

	fmt.Printf("loops: %v\n", loops)
}

type grid struct {
	fields           [][]field
	originX, originY int
}

func (g grid) reset() {
	for y, line := range g.fields {
		for x, field := range line {
			if field == visited || field == customObstacle || field == guard {
				g.fields[y][x] = empty
			}
		}
	}
	g.fields[g.originY][g.originX] = guard
}

type field int

const (
	empty field = iota
	obstacle
	customObstacle
	guard
	visited
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

func readInput(path string) (grid, error) {
	file, err := os.Open(path)
	if err != nil {
		return grid{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	g := grid{fields: make([][]field, 0), originX: -1, originY: -1}

	currentLine := -1
	for scanner.Scan() {
		currentLine++
		line := scanner.Text()
		g.fields = append(g.fields, make([]field, len(line)))
		for currentField, field := range line {
			switch field {
			case '.':
				g.fields[currentLine][currentField] = empty
			case '#':
				g.fields[currentLine][currentField] = obstacle
			case '^':
				g.fields[currentLine][currentField] = guard
				g.originX = currentField
				g.originY = currentLine
			default:
				return grid{}, fmt.Errorf("Unknown field type: %v", field)
			}
		}
	}

	return g, nil
}

func (g grid) printGrid(dir direction) {
	for _, line := range g.fields {
		for _, field := range line {
			switch field {
			case empty:
				fmt.Print("\033[0;0m.")
			case obstacle:
				fmt.Print("\033[0;31m#")
			case customObstacle:
				fmt.Print("\033[0;32mO")
			case guard:
				char := '^'
				switch dir {
				case up:
					char = '^'
				case right:
					char = '>'
				case down:
					char = 'v'
				case left:
					char = '<'
				}
				fmt.Print("\033[0;34m" + string(char))
			case visited:
				fmt.Print("\033[0;34mX")
			}
		}
		fmt.Println()
	}
}

func (g grid) run() ([][]int, error) {
	guardX, guardY, err := g.findGuard()
	if err != nil {
		return nil, err
	}

	path := make([][]int, 0)
	hits := make([]hit, 0)

	guardDir := up
	for true {
		nextFieldX, nextFieldY, err := g.getNextField(guardX, guardY, guardDir)
		if err != nil {
			g.fields[guardY][guardX] = visited
			break
		}

		switch g.fields[nextFieldY][nextFieldX] {
		case empty, visited:
			g.fields[guardY][guardX] = visited
			path = append(path, []int{guardX, guardY})
			guardX, guardY = nextFieldX, nextFieldY
			g.fields[guardY][guardX] = guard
		case obstacle, customObstacle:
			for _, hit := range hits {
				if hit.x == nextFieldX && hit.y == nextFieldY && hit.dir == guardDir {
					return nil, fmt.Errorf("Loop detected")
				}
			}
			hits = append(hits, hit{x: nextFieldX, y: nextFieldY, dir: guardDir})
			guardDir = turnRight(guardDir)
		}
	}

	return path, nil
}

type hit struct {
	x, y int
	dir  direction
}

func (g grid) getNextField(x, y int, dir direction) (int, int, error) {
	x, y = applyDirection(x, y, dir)

	if x < 0 || y < 0 || y >= len(g.fields) || x >= len(g.fields[y]) {
		return 0, 0, fmt.Errorf("Out of bounds")
	}

	return x, y, nil
}

func applyDirection(x, y int, dir direction) (int, int) {
	switch dir {
	case up:
		y--
		break
	case right:
		x++
		break
	case down:
		y++
		break
	case left:
		x--
		break
	}

	return x, y
}

func turnRight(dir direction) direction {
	switch dir {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	}
	return dir
}

func (g grid) findGuard() (int, int, error) {
	for y, line := range g.fields {
		for x, field := range line {
			if field == guard {
				return x, y, nil
			}
		}
	}

	return 0, 0, fmt.Errorf("Guard not found")
}

func (g grid) countVisitedFields() int {
	count := 0
	for _, line := range g.fields {
		for _, field := range line {
			if field == visited {
				count++
			}
		}
	}
	return count
}

func (g grid) countLoops() (int, error) {
	g.reset()
	path, err := g.run()
	if err != nil {
		return -1, err
	}

	pathWithoutOrigin := make([][]int, 0)
	for _, field := range path {
		if field[0] != g.originX || field[1] != g.originY {
			pathWithoutOrigin = append(pathWithoutOrigin, []int{field[0], field[1]})
		}
	}

	uniquePaths := make([][]int, 0)
	for _, field := range pathWithoutOrigin {
		alreadyIn := false
		for _, uniqueField := range uniquePaths {
			if field[0] == uniqueField[0] && field[1] == uniqueField[1] {
				alreadyIn = true
			}
		}
		if !alreadyIn {
			uniquePaths = append(uniquePaths, field)
		}
	}

	loops := 0
	for _, field := range uniquePaths {
		g.reset()
		g.fields[field[1]][field[0]] = customObstacle
		_, err := g.run()
		if err != nil && err.Error() == "Loop detected" {
			loops++
		}
	}

	return loops, nil
}
