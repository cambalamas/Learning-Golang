package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	sleepTime = 100 * time.Millisecond
)

var (
	sentences = os.Args[1:]
)

func main() {
	for _, s := range sentences {
		slowPrint(s + "\n")
	}
}

func slowPrint(s string) {
	stringArray := strings.Split(s, "")
	for _, char := range stringArray {
		fmt.Print(char)
		time.Sleep(sleepTime)
	}
}
