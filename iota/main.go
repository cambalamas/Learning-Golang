package main

import (
	"fmt"
	"time"
)

const (
	mon = time.Monday * iota
	tue
	wed
	thu
	fri
	sat
	sun
)

const (
	jan = time.January + iota
	feb
	mar
	apr
	may
	jun
	jul
	aug
	sep
	oct
	nov
	dec
)

func main() {
	fmt.Println(mon, tue, wed, thu, fri, sat, sun)
	fmt.Println(jan, feb, mar, apr, may, jun, jul, aug, sep, oct, nov, dec)
}
