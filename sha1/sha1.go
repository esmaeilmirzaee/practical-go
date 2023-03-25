package main

import (
	"compress/gzip"
	"crypto/sha512"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	sig, err := sha1("github_2.tar.gz")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	log.Println(sig)
}

func sha1(fileName string) (string, error) {
	/*
		idiom: acquire a resource, check for errors, defer a release
	*/
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	r, err := gzip.NewReader(file)
	if err != nil {
		return "", err
	}

	w := sha512.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := sha512.Sum512(nil)

	return fmt.Sprintf("%x", sig), nil
}
