package main

import "fmt"

func main() {
	/*
		1. After the goroutine starts there is no way to access it
		2. Defer doesn't wait for all the goroutines

	*/
	go fmt.Println("Goroutines")
	fmt.Println("main")
}
