package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func chker(err error) {
	if err != nil {
		log.Fatalf("[!]: %v", err)
	}
}

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		chker(err)
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
