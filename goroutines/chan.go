package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		1. After the goroutine starts there is no way to access it
		2. Defer doesn't wait for all the goroutines
	*/
	go fmt.Println("Goroutines")
	fmt.Println("main")

	for i := 0; i < 3; i++ {
		// Fix 1: pass the parameter
		go func(n int) {
			fmt.Printf("i is %d \t", n+1)
		}(i)
		// Fix 2: use a shadow variable
		i := i
		go func() {
			fmt.Printf("i2 is %d\t", i)
		}()
		/* BUG: All the goroutines fires with the same `i`;
		this's a closure so each environment has the same variables
		closure means an environment that remembers the variables which has been created with.
		go func() {
			fmt.Printf("i: %d\t", i)
		}()
		*/
	}
	fmt.Println()
	time.Sleep(time.Second)
}
