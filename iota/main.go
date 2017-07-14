package main

import (
	"fmt"
	"time"
)

const (
	monday = time.Monday * iota
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

func main() {
	fmt.Print(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
}
