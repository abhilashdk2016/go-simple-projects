package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// what is the output
	// fix the issue.

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i) // Pass i to go-routine so that closure comes into picture and go-routine has the ability to remember value of i
	}
	wg.Wait()
}
