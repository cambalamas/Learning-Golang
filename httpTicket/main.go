package main

import (
	"fmt"
	"log"
	"net/http"
)

var ticket = make(chan int)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "<h1>Ticket number: %v</h1>", <-ticket)
	if err != nil {
		log.Fatalf("Error on get ticket number: %v", err)
	}
}

func main() {
	http.HandleFunc("/next", rootHandler)
	go func() {
		for i := 0; ; i++ {
			ticket <- i
		}
	}()
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
