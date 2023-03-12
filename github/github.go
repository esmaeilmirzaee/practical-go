package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"unicode/utf8"
)

func main() {
	g := "https://api.github.com/users/esmaeilmirzaee"
	resp, err := http.Get(g)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("%s", resp.Status)
	}
	//fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	//if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
	//	log.Fatalf("error: can't copy %s", err)
	//}

	reply, err := decodeReply(resp.Body)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	fmt.Println("DECODE REPLY", reply)
}

func isPalindrome(s string) bool {
	width := utf8.RuneCountInString(s) - 1
	for i := 0; i < width/2; i++ {
		if s[i] != s[width-1-i] {
			return false
		}
	}
	return true
}

/**
 * Get name and number of repo
 */
type Reply struct {
	Name        string
	PublicRepos int `json:"public_repos"`
}

func decodeReply(s io.Reader) (Reply, error) {
	var r Reply
	dec := json.NewDecoder(s)
	err := dec.Decode(&r)

	return r, err
}

func githubInfo(name string) (Reply, error) {
	g := fmt.Sprintf("https://api.github.com/%s", name)
	resp, err := http.Get(g)
	if err != nil {
		log.Fatalf("%s", err)
	}

	dec := json.NewDecoder(resp.Body)
	var r Reply
	err = dec.Decode(&r)

	return r, err
}
