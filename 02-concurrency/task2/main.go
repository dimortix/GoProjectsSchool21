package main

import (
	"flag"
	"fmt"
)

func main() {
	kFlag := flag.Int("K", -1, "generator start")
	nFlag := flag.Int("N", -1, "generator end")
	flag.Parse()

	if (*kFlag < 0 || *nFlag < 0) && len(flag.Args()) >= 2 {
		_, _ = fmt.Sscanf(flag.Args()[0], "%d", kFlag)
		_, _ = fmt.Sscanf(flag.Args()[1], "%d", nFlag)
	}

	k, n := *kFlag, *nFlag
	if k == -1 || n == -1 {
		return
	}

	for v := range square(generator(k, n)) {
		fmt.Println(v)
	}
}

func generator(k, n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		if k <= n {
			for i := k; i <= n; i++ {
				out <- i
			}
			return
		}
		for i := k; i >= n; i-- {
			out <- i
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}()
	return out
}
