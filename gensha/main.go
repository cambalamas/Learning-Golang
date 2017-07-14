package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args)-1 != 1 {
		log.Fatal("USAGE: ./shasum <alphanum>")
		os.Exit(2)
	}

	ss := sha256.Sum256([]byte(os.Args[1]))
	if _, err := fmt.Fprintf(os.Stdout, "%x\n", ss); err != nil {
		log.Println("Could not write into Stdout")
	}
}
