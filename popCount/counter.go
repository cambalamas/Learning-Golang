package main

var stack [256]byte

func init() {
	for i := range stack {
		stack[i] = stack[i/2] + byte(i&1)
	}
}

// PopCount returns the population count of x (as number of set bits)
func PopCount(x uint64) int {

	var i uint64
	var count byte
	for ; i < x; i++ {
		count += stack[byte(x>>(uint(i)*8))]
	}

	return int(count)
}
