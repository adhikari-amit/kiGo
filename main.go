package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	// A goroutine is a lightweight thread managed by the Go runtime
	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"phi",
		"che",
		"zeta",
	}

	for _, word := range words {
		// Init a thred by using "go" keyword infront of the command
		go fmt.Println(word)
	}

	// Introducting a 1 sec sleep just to wait for returning the result from the threads
	time.Sleep(1 * time.Second)

	printSomething("amamit")
	// Observe the output. This is due to not managing the threads we created in the go

}
