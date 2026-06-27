package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	wordsLine, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	wordsLine = strings.TrimSpace(wordsLine)

	kLine, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	kLine = strings.TrimSpace(kLine)

	if wordsLine == "" {
		return
	}

	k, err := strconv.Atoi(kLine)
	if err != nil || k <= 0 {
		return
	}

	words := strings.Fields(wordsLine)
	top := TopKFrequent(words, k)
	if len(top) == 0 {
		return
	}

	fmt.Println(strings.Join(top, " "))
}

func TopKFrequent(words []string, k int) []string {
	if len(words) == 0 || k <= 0 {
		return []string{}
	}

	frequency := make(map[string]int, len(words))
	for _, w := range words {
		frequency[w]++
	}

	uniqueWords := make([]string, 0, len(frequency))
	for w := range frequency {
		uniqueWords = append(uniqueWords, w)
	}

	sort.Slice(uniqueWords, func(i, j int) bool {
		wi, wj := uniqueWords[i], uniqueWords[j]
		fi, fj := frequency[wi], frequency[wj]
		if fi == fj {
			return wi < wj
		}
		return fi > fj
	})

	if k > len(uniqueWords) {
		k = len(uniqueWords)
	}

	result := make([]string, k)
	copy(result, uniqueWords[:k])

	return result
}

