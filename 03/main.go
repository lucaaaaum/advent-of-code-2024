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

	multiply(input)
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

func multiply(input string) int {
	for i, v := range input {
		if i > len(input)-8 {
			break
		}

		if v == 'm' {
			fmt.Printf("input[i : i+8]: %v\n", input[i:i+8])
			if input[i:i+4] != "mul(" {
                continue
			}


		}
	}

	return 1
}
