package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var exeName = filepath.Base(os.Args[0])

func main() {

	if len(os.Args)-1 == 0 {
		if _, err := fmt.Fprintf(os.Stderr, "Usage: %s <string1> <string2>...", exeName); err != nil {
			log.Fatalln(err)
		}
		os.Exit(2)
	}

	sentences := os.Args[1:]
	for _, s := range sentences {
		PrintLikeRpg(s + "\n")
	}
}
