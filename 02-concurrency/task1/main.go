package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type pair struct {
	idx int
	ms  int
}

func main() {
	nFlag := flag.Int("N", -1, "number of goroutines")
	mFlag := flag.Int("M", -1, "max sleep time in milliseconds")
	flag.Parse()

	rest := flag.Args()
	if (*nFlag < 0 || *mFlag < 0) && len(rest) >= 2 {
		var n, m int
		if _, err := fmt.Sscanf(rest[0], "%d", &n); err == nil {
			*nFlag = n
		}
		if _, err := fmt.Sscanf(rest[1], "%d", &m); err == nil {
			*mFlag = m
		}
	}

	if *nFlag <= 0 {
		return
	}
	if *mFlag < 0 {
		*mFlag = 0
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sleepMs := make([]int, *nFlag)
	for i := 0; i < *nFlag; i++ {
		sleepMs[i] = r.Intn(*mFlag + 1)
	}

	results := make([]pair, *nFlag)
	var wg sync.WaitGroup
	wg.Add(*nFlag)

	for i := 0; i < *nFlag; i++ {
		i := i
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(sleepMs[i]) * time.Millisecond)
			results[i] = pair{idx: i, ms: sleepMs[i]}
		}()
	}

	wg.Wait()

	sort.Slice(results, func(i, j int) bool {
		if results[i].ms == results[j].ms {
			return results[i].idx < results[j].idx
		}
		return results[i].ms > results[j].ms
	})

	for _, p := range results {
		fmt.Printf("%d %d\n", p.idx, p.ms)
	}
}
