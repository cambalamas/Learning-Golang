package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	log.SetOutput(os.Stderr)

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatalln(err)
		os.Exit(2)
	}

	for _, i := range interfaces {
		if strings.Contains(i.Flags.String(), "up") {
			fmt.Printf("[%s]\t%q\n", "UP", i.Name)
		} else {
			fmt.Printf("[%s]\t%q\n", "DOWN", i.Name)
		}
	}
}
