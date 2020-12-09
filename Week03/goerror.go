package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

// @Project: Go-000
// @Author: houseme
// @Description:
// @File: goerror
// @Version: 1.0.0
// @Date: 2020/12/9 21:07
// @Package Week03

func main() {
	g := new(errgroup.Group)
	var urls = []string{
		"http://www.loudijie.org/",
		"http://www.yuerso.com/",
		"http://www.baidu.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}
}
