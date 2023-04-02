package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

	wordFreq(file)
}

func wordFreq(r io.Reader) (map[string]int, error) {
	//	hashMap := make(map[string]int)
	s := bufio.NewScanner(r)
	lnums := 0
	for s.Scan() {
		lnums++
		l := s.Text() // Current line
		for i, w := range l {
			if i == 10 {
				break
			}
			fmt.Println(string(w))
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	fmt.Printf("Number of lines: %d", lnums)
	return nil, nil
}
