package concat

import (
	"fmt"
	"testing"
)

var slices = [][]byte{
	[]byte("my first slice"),
	[]byte("second slice"),
	[]byte("third slice"),
	[]byte("fourth slice"),
	[]byte("fifth slice"),
}

var B []byte

func BenchmarkConcatCopyPreAllocate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		B = concatCopyPreAllocate(slices)
		fmt.Print(B)
	}
}

func BenchmarkConcatAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		B = concatAppend(slices)
		fmt.Print(B)
	}
}
