package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Create a pair list (e.g.: counts['hi']=5)
	counts := make(map[string]int)
	// Create a way to read user input
	input := bufio.NewScanner(os.Stdin)

	// Write the document
	fmt.Print("Write, (^Z to stop):\n")
	for input.Scan() {
		/**
		* input.Text() is what the user typed
		* so if him type more than once the same word
		* the value will increment.
		*/
		counts[input.Text()]++
	}

	// Print pair line and ocurrences
	for line, n := range counts {
		fmt.Printf("%s: %d\n", line, n)
	}
}
