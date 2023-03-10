package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s = "Go 😃"
	banner(s, 6)
	// byte (uint8)
	// rune (int32)
	for idx, char := range s {
		fmt.Printf("\n %d: %c %T", idx, char, char)
	}

	fmt.Println(isPalindrome("gogog"))
}

func banner(text string, width int) {
	padding := (width - utf8.RuneCountInString(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Print(text)
	fmt.Println()
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}

}

func isPalindrome(text string) bool {
	width := utf8.RuneCountInString(text) - 1
	fmt.Println()
	fmt.Println()
	for i := 0; i < width/2; i++ {
		if text[i] != text[width-i] {
			return false
		}
	}
	return true
}
