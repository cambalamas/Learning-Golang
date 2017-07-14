// Select, permite la espera de varios canales.

package main

import (
	"fmt"
	"time"
)

var (
	c1   = make(chan string)
	c2   = make(chan string)
	quit = make(chan string)
)

func main() {

	// Comunicaciones con el primer canal
	go func() {
		c1 <- "one"
		time.Sleep(time.Second * 3)
		c1 <- "one again"
	}()

	// Comunicaciones con el segundo canal
	go func() {
		time.Sleep(time.Second * 4)
		c2 <- "two"
	}()

	// Comunicaciones con el canal de salida
	go func() {
		time.Sleep(5 * time.Second)
		quit <- ""
	}()

loop:
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case <-quit:
			break loop
		}
	}
}
