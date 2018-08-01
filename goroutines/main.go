package main

import (
	"fmt"
	"sync"
)

func count(prefix string, count int, wg *sync.WaitGroup) {
	// Defer the waitgroup Done
	defer wg.Done()

	for i := 0; i < count; i++ {
		// Print out the prefix and the loop variable `i`
		fmt.Printf("Prefix: %s, i=%v", prefix, i)
	}
}

func main() {
	// Define the wait group
	wg := &sync.WaitGroup{}

	// Increment wait group
	wg.Add(2)

	// Call count with a goroutine, first argument "first", second argument 50,
	// third argument the wait group
	go count("first", 50, wg)
	// Call count with a goroutine, first argument "second", second argument 50,
	// third argument the wait group
	go count("second", 50, wg)

	// Wait...
	wg.Wait()
}
