package main

import (
	"fmt"
	"time"
)

func sleepAndTalk(secs time.Duration, msg string, c chan string) {
	time.Sleep(secs * time.Second)
	c <- msg
}

func main() {
	c := make(chan string)

	go sleepAndTalk(0, "Hola,", c)
	go sleepAndTalk(1, "¿Cómo", c)
	go sleepAndTalk(2, "estan", c)
	go sleepAndTalk(3, "ustedes?", c)

	for i := 0; i < 4; i++ {
		fmt.Printf("%v ", <-c)
	}

}
