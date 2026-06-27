package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	left := readFloat64(reader, "Input left operand:")
	op := readOperation(reader, "Input operation:")
	right := readFloat64(reader, "Input right operand:")

	var result float64

	switch op {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		for right == 0 {
			fmt.Println("Invalid input")
			right = readFloat64(reader, "Input right operand:")
		}
		result = left / right
		fmt.Printf("%.3f\n", result)
		return
	default:
		fmt.Println("Invalid input")
		return
	}

	if result == float64(int64(result)) {
		fmt.Printf("Result: ")
		fmt.Printf("%d\n", int64(result))
	} else {
		fmt.Printf("%g\n", result)
	}
}

func readFloat64(reader *bufio.Reader, prompt string) float64 {
	for {
		fmt.Println(prompt)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		line = strings.TrimSpace(line)
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		return value
	}
}

func readOperation(reader *bufio.Reader, prompt string) string {
	for {
		fmt.Println(prompt)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		line = strings.TrimSpace(line)
		switch line {
		case "+", "-", "*", "/":
			return line
		default:
			fmt.Println("Invalid input")
		}
	}
}
