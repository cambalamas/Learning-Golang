package main

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
)

const friday = time.Friday

var (
	msg   string
	today = time.Now().Weekday()
)

func main() {
	msg = "Is friday?\n"

	if is, left := isFriday(); is {
		msg += "YES!"
	} else {
		msg += fmt.Sprintf("NO! Left %v days", left)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "%v", msg); err != nil {
			logrus.Warn(err)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		logrus.Fatal(err)
	}
}

// isFriday give the truth of friday or left days to reach on it
func isFriday() (is bool, left float64) {
	is = today == friday
	left = math.Abs(float64(today - friday))
	return
}
