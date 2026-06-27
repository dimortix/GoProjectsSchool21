package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Input first list of numbers:")
	reader := bufio.NewReader(os.Stdin)

	firstLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	fmt.Println("Input second list of numbers:")

	secondLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	fmt.Println("Result:")
	firstLine = strings.TrimSpace(firstLine)
	secondLine = strings.TrimSpace(secondLine)

	firstSlice, ok := parseIntSlice(firstLine)
	if !ok {
		fmt.Println("Invalid input")
		return
	}
	secondSlice, ok := parseIntSlice(secondLine)
	if !ok {
		fmt.Println("Invalid input")
		return
	}

	intersection := intersectUnique(firstSlice, secondSlice)
	if len(intersection) == 0 {
		fmt.Println("Empty intersection")
		return
	}

	values := make([]string, len(intersection))
	for i, v := range intersection {
		values[i] = strconv.Itoa(v)
	}

	fmt.Println(strings.Join(values, " "))
}

func parseIntSlice(line string) ([]int, bool) {
	if strings.TrimSpace(line) == "" {
		return []int{}, true
	}

	parts := strings.Fields(line)
	result := make([]int, len(parts))
	for i, p := range parts {
		value, err := strconv.Atoi(p)
		if err != nil {
			return nil, false
		}
		result[i] = value
	}
	return result, true
}

func intersectUnique(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return []int{}
	}

	inSecond := make(map[int]struct{}, len(b))
	for _, v := range b {
		inSecond[v] = struct{}{}
	}

	seen := make(map[int]struct{})
	result := make([]int, 0)

	for _, v := range a {
		if _, ok := inSecond[v]; ok {
			if _, already := seen[v]; !already {
				seen[v] = struct{}{}
				result = append(result, v)
			}
		}
	}

	return result
}
