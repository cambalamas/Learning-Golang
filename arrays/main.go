package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	// Dar la vuelta a un slice
	sort.Stable(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)
	// Borrar un elemento intermedio respetando el orden
	arr = append(arr[:2], arr[3:]...)
	fmt.Println(arr)

	dayStr := make(map[int]string)
	// Populate the map
	for sec := 0; sec < 60; sec++ {
		dayStr[sec] = "Here must be a cool sentence!"
	}

}

// UReadMap, reads a map on default map order (usually unordered)
func UReadMap(m map[int]string) {
	for k, v := range m {
		fmt.Printf("[%v sec]: %v\n", k, v)
	}
}

// OReadMap, reads a map in ordered way
func OReadMap(m map[int]string) {
	for i := 0; i < 60; i++ {
		fmt.Printf("[%v sec]: %v\n", i, dayStr[i])
	}
}
