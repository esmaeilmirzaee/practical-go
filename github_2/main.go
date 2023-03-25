package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	url2 "net/url"
)

func main() {
	fmt.Println(githubInfo("esmaeilmirzaee"))
}

type Reply struct {
	Name          string
	NumberOfRepos int `json:"public_repos"`
}

func githubInfo(username string) (string, int, error) {
	url := "https://api.github.com/users/" + url2.PathEscape(username)
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%s", resp.Status)
	}
	var r Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}

	return r.Name, r.NumberOfRepos, nil
}
