package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Sirupsen/logrus"
)

var err error
var cant uint64

func main() {

	if len(os.Args) != 2 {
		logrus.Fatal("USAGE: ./popCount <uint>")
		os.Exit(1)
	}

	if cant, err = strconv.ParseUint(os.Args[1], 10, 64); err != nil {
		logrus.Fatal("USAGE: ./popCount <uint>")
		os.Exit(1)
	}

	fmt.Println(PopCount(cant))
}
