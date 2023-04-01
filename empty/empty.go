package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func main() {
	var i any
	// go < 1.18
	// var i interface{}

	i = 7
	fmt.Println(i)

	i = "hi"
	fmt.Println(i)

	// comma, ok
	n, ok := i.(int)
	if ok {
		fmt.Println(n)
	} else {
		_ = fmt.Errorf("error: cann't conver. %d", i)
	}
	// switch type
	switch i.(type) {
	case int:
		fmt.Println("an int")
	case string:
		fmt.Println("a string")
	default:
		fmt.Printf("%T", i)
	}

	fmt.Println("Max integer", maxFloat64([]float64{1, 2, 4, 5, 2, 3, 8}))
	fmt.Println("Max integer", maxInt([]int{1, 2, 4, 5, 2, 3, 8}))
	fmt.Println("Max integer", max([]int{1, 2, 4, 5, 2, 3, 8}))
	keys := []key{1, 2, 6, 3, 4, 54}
	if slices.Contains(keys, 2) {
		fmt.Printf("%d exists in keys", 2)
	}
}

type key byte

func containsKey(keys []key, k key) bool {
	for _, k2 := range keys {
		if k == k2 {
			return true
		}
	}
	return false
}

type Number interface {
	int | float64
}

func max[T Number](nums []T) T {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}

	return max
}

func maxInt(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}

	return max
}

func maxFloat64(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}

	return max
}
