package main

import (
	"fmt"
)

func process(messages chan string, quit chan struct{}) {
	// Loop through your messages
	for m := range messages {
		// print the message with a for loop using range
		fmt.Println(m)
	}
	close(quit)
}

func main() {
	// declare the messages channel of type string and capacity of 5
	msgCh := make(chan string, 5)

	// declare a signal channel
	quitCh := make(chan struct{})

	// launch process in a goroutine
	go process(msgCh, quitCh)

	// declare 5 fruits in a []string
	fruits := []string{"apple", "orange", "pear", "another apple", "annother orange"}

	// loop through the fruits and send them to the messages channel
	for _, f := range fruits {
		msgCh <- f
	}

	// close the messages channel

	close(msgCh)
	// wait for everything to finish

	<- quitCh

	fmt.Println("done")
}