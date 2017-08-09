package main

import (
	"flag"
	"fmt"
	"time"
)

var stretchInterval = flag.Duration("interval", 10*time.Minute, "stretch interval")

func main() {
	for {
		fmt.Print("\x1b[2J\x1b[1;1H\x07")
		fmt.Println("\x1b[1mTime to stretch!\x1b[0m")
		fmt.Scanln()

		fmt.Print("\x1b[2J\x1b[1;1H")
		time.Sleep(*stretchInterval)
	}
}
