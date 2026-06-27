package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	kFlag := flag.Uint("K", 0, "ticker interval in seconds")
	flag.Parse()

	if *kFlag == 0 && len(flag.Args()) >= 1 {
		_, _ = fmt.Sscanf(flag.Args()[0], "%d", kFlag)
	}
	if *kFlag == 0 {
		return
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigCh)

	var (
		tick    uint = 1
		elapsed uint = 0
	)

	for {
		for s := uint(0); s < *kFlag; s++ {
			select {
			case <-sigCh:
				fmt.Println("Termination")
				return
			default:
			}
			time.Sleep(time.Second)
		}

		elapsed += *kFlag
		fmt.Printf("Tick %d since %d\n", tick, elapsed)
		tick++
	}
}
