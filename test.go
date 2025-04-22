package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond) // Give other goroutines a chance to run
	}
}

func main() {
	// Run say("World") as a goroutine
	go say("World")

	// Run say("Hello") in the main goroutine (sequentially)
	say("Hello")

	fmt.Println("Main function finished.")

	// --- Problem Point (explained below) ---
	// If we didn't have the say("Hello") call which takes time,
	// or some other way to wait, the main function might exit
	// before the "World" goroutine gets a chance to run properly.
}
