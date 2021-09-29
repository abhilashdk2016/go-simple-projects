package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct Call
	fun("Direct Call")

	// goroutine function call
	go fun("goroutine function call")

	// goroutine with anonymous function
	go func() {
		fun("goroutine with anonymous function")
	}()

	// goroutine with function value call
	fv := fun

	go fv("goroutine with function value call")

	// wait for goroutines to end
	fmt.Println("Waiting for goroutine to complete...")
	time.Sleep(1 * time.Second)

	fmt.Println("done...")
}
