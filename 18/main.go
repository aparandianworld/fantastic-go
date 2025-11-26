package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func checkURL(url string) bool {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // handle timeout with context

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return false
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true
	}

	return false

}

func main() {
	gitURL := "https://github.com/aparandianworld"
	fmt.Printf("Status: %+v\n", checkURL(gitURL))
}
