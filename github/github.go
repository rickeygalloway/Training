/*
Write a function
	githubInfo(login string) (string, int, err)
Which returns the name of the user and number of public repositories

$ curl https://api.github.com/users/shipt
*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	name, numRepos, err := githubInfo("shipt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s, public repos: %d\n", name, numRepos)
}

func githubInfo(login string) (string, int, error) {

	u := fmt.Sprintf("https://api.github.com/users/%s", url.PathEscape(login))
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return "", 0, err
	}

	// resp, err := http.Get(u)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("status: %s", resp.Status)
	}
	defer resp.Body.Close()

	//var ur userResponse
	var ur struct { // anonymous struct
		Name     string
		NumRepos int `json:"public_repos"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&ur); err != nil {
		return "", 0, err
	}

	return ur.Name, ur.NumRepos, nil
}
