package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
Some functions run before the main function
1. some variables declaration because they might be
2. init function
func init() {}
*/
func main() {
	n, err := os.Getwd()
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	log.Println("PATH: ", n)

	file, err := os.Open("./assets/sherlock.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

			log.Fatalf("error: %s; couldn't close the file.", err)
		}
	}(file)
	/*
		backticks create raw strings so
		\ is just \
		and help to create multi-line strings
	*/
	displayMax(file)
}

func displayMax(r io.Reader) {
	freqs, err := wordFreq(r)
	if err != nil {
		log.Fatalf("error on word frequency: %s\n", err)
	}

	word, count, err := maxWord(freqs)
	if err != nil {
		log.Fatalf("error on finding the maximum: %s\n", err)
	}

	log.Printf("The number of occurance of %s is %d\n", word, count)
}

// 's or 't
var apostropheReg = regexp.MustCompile(`'s | 't`)

func removeApostropheS(r io.Reader) io.Reader {
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := s.Text()
		remove := apostropheReg.ReplaceAll([]byte(l), []byte(""))
		fmt.Println(string(remove))
	}
	return nil
}

// find all words
var wordRegx = regexp.MustCompile(`[a-zA-Z]+`)

func wordFreq(r io.Reader) (map[string]int, error) {
	//hashMap := map[string]int // [word]: count; also this is a nil map, in other words it's uninitialized
	/*
		hashMap := make(map[string]int) // initializes the hashMap
		or
		hashMap := map[string]int {
			"appl": 173.23,
		}
	*/
	s := bufio.NewScanner(r)
	hashMap := make(map[string]int) // [word]: count
	lnums := 0
	for s.Scan() {
		l := s.Text() // current line
		words := wordRegx.FindAllString(l, -1)
		for _, w := range words {
			hashMap[strings.ToLower(w)]++
		}
		lnums++
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	fmt.Printf("Number of lines: %d\n", lnums)
	return hashMap, nil
}

func maxWord(words map[string]int) (string, int, error) {
	maxW, maxCount := "", 0

	for w, _ := range words {
		if words[w] > maxCount {
			//fmt.Printf("w: %s, count: %d, maxw: %s, maxCount: %d", w, words[w], maxW, maxCount)
			maxW, maxCount = w, words[w]
		}
	}

	return maxW, maxCount, nil
}
