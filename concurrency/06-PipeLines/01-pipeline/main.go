package main

import "fmt"

// generator() => square() => print

// converts list of integers to a channel
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// receive on inbound channel
// square the number
// output to outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// set up pipeline
	for n := range square(generator(2, 3)) {
		fmt.Println(n)
	}
}
