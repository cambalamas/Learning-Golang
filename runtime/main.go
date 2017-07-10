package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("\n[%v]: %v\n", "ARCH", runtime.GOARCH)
	fmt.Printf("[%v]: %v\n\n", "OS", runtime.GOOS)
	fmt.Printf("[%v]: %v\n", "NUMCPU", runtime.NumCPU())
	fmt.Printf("[%v]: %v\n", "GOMAXPROCS", runtime.GOMAXPROCS(0))
	fmt.Printf("[%v]: %v\n", "VERSION", runtime.Version())
	fmt.Printf("[%v]: %v\n\n", "GOROOT", runtime.GOROOT())
}
