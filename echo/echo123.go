package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Vars declaration
	var s, sep string

	// 1) Clasic loop. (Ugly)
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("For: " + s)

	// 2) Foreach loop. (Practical)
	s = ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("Foreach: " + s)

	// 3) Join them (Efficient)
	s = strings.Join(os.Args[1:], " ")
	fmt.Println("Strings Join: " + s)
}
