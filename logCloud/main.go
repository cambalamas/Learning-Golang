package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/signal"
)

func main() {
	go rwText()
	signalsHandle()
}

func rwText() {
	sc := bufio.NewScanner(io.TeeReader(os.Stdin, os.Stdout))
	for {
		sc.Scan()
	}
}

// es una patata
func signalsHandle() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	logFile, err := os.Create("signals.log")
	if err != nil {
		log.SetOutput(os.Stderr)
		log.Fatalf("Could not create the file %q: %q", logFile.Name(), err)
	}

	for {
		s := <-signals
		if s != nil {
			log.SetOutput(logFile)
			log.Fatalf("Terminated by: %q", s)
		}
	}
}

// func errorf(errNum int, format string, args ...interface{}) {
// 	log.SetOutput(os.Stderr)
// 	log.Fatalf(format, args)
// 	if errNum != 0 {
// 		os.Exit(errNum)
// 	}
// }
