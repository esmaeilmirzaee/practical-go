package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Printf("%d/%d=%d", 2, 9, safeDiv(2, 9))
	q, err := safeDiv(2, 0)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%d/%d=%d", 2, 9, q)
}

func safeDiv(a, b int) (q int, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("error: %v", e)
			err = fmt.Errorf("error: %v", e)
		}
	}()

	// return a / b, nil
	q = a / b
	return

}

// unsafe division, b == 0 panics
func div(a, b int) int {

	return a / b
}
