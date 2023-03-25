package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(isPalindrome("gog"))
	fmt.Println(isPalindrome("g✅g"))
	fmt.Println(isPalindrome("gogo"))

	fmt.Println(isPalindrome2("gog"))
	fmt.Println(isPalindrome2("g✅g"))
	fmt.Println(isPalindrome2("gogo"))
}

func isPalindrome(s string) bool {
	length := utf8.RuneCountInString(s) / 2
	lowered_s := strings.ToLower(s)

	for i := 0; i < length; i++ {
		if lowered_s[i] != lowered_s[length-i-1] {
			return false
		}
	}
	return true
}

func isPalindrome2(s string) bool {
	rune_s := []rune(s)
	for i := 0; i < (len(rune_s)/2)-i-1; i++ {
		if rune_s[i] != rune_s[(len(rune_s)/2)-i-1] {
			return false
		}
	}
	return true
}
