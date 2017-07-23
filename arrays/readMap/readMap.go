package readMap

import "fmt"

// UReadMap reads a map on default map order (usually unordered)
func UReadMap(m map[int]string) bool {
	for k, v := range m {
		fmt.Printf("[%v sec]: %v\n", k, v)
	}
	return true
}

// OReadMap reads a map in ordered way
func OReadMap(m map[int]string) bool {
	for i := 0; i < 60; i++ {
		fmt.Printf("[%v sec]: %v\n", i, m[i])
	}
	return true
}
