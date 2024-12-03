package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    input, err := readInput()
    if err != nil {
        panic(err)
    }

    fmt.Printf("input: %v\n", input)
}

func readInput() (string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)

    input := ""

    for scanner.Scan() {
        input += scanner.Text()
    }

    return input, nil
}
