package main

import (
	"fmt"
)

func main() {
	//fmt.Println(concat([]string{"A", "B"}, []string{"C", "D", "E"}))
	//someFundamentals()
}

func someFundamentals() {
	var a []int
	if a == nil {
		fmt.Println("nil slice")
	}

	a2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(a2)
	var a4 []int
	for i := 0; i < 10; i++ {
		a4 = appendIntSlice(a4, i)
	}
	printSlice(a4)

	a3 := a2[1:4]
	printSlice(a3)

	a3 = append(a3, 200, 100)
	printSlice(a3)
}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d, %#v\n", len(s), cap(s), s)
}

func appendIntSlice(s []int, newValue int) []int {
	i := len(s)

	printSlice(s)
	if len(s) < cap(s) {
		s = s[:len(s)+1]
	} else {
		s2 := make([]int, 2*len(s)+1)
		fmt.Printf("Reallocate: %d to %d\n", i, len(s2))
		copy(s2, s)
		s = s2[:len(s)+1]
	}

	fmt.Printf("Where to store %d\n", i)
	printSlice(s)
	s[i] = newValue
	return s
}

func concat(s1, s2 []string) []string {
	// restrictions no for loops
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s
}

func concat2(s1, s2 []string) []string {
	return append(s1, s2...)
}
