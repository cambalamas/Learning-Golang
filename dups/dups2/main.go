package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func countOnFile(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func countOnStdin(f io.Reader, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Print("Write, (^Z to stop):\n")
		countOnStdin(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				_, err1 := fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				// Writer error check
				if err1 != nil {
					log.Print("Stderr failed")
				}
				// Follow the execution after notify it
				continue
			}
			countOnFile(f, counts)
			err2 := f.Close()
			// Close error check
			if err2 != nil {
				log.Fatal("Can´t close :(")
			}
		}
	}
	for line, n := range counts {
		fmt.Printf("%s: %d\n", line, n)
	}
}
