package readMap

import (
	"testing"
)

var dayStr = make(map[int]string)

func init() {
	for sec := 0; sec < 60; sec++ {
		dayStr[sec] = "Here must be a cool sentence!"
	}
}

func TestUReadMap(t *testing.T) {
	if !UReadMap(dayStr) {
		t.Error("Can't read the map in unordered way")
	}
}
func TestOReadMap(t *testing.T) {
	if !OReadMap(dayStr) {
		t.Error("Can't read the map in ordered way")
	}
}

func BenchmarkUReadMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UReadMap(dayStr)
	}
}

func BenchmarkOReadMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		OReadMap(dayStr)
	}
}
