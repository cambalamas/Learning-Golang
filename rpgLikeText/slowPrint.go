package main

import (
	"fmt"
	"time"
)

const sleepTime = 100 * time.Millisecond

// PrintLikeRpg print a string with a delay between chars.
func PrintLikeRpg(s string) {
	for _, char := range s {
		fmt.Print(string(char))
		time.Sleep(sleepTime)
	}
}
