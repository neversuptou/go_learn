package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

func ping(url string, wg *sync.WaitGroup, respCh chan int, errCh chan error) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		errCh <- err
		return
	}
	defer resp.Body.Close()

	respCh <- resp.StatusCode
}

func main() {
	wg := sync.WaitGroup{}
	path := flag.String("file", "url.txt", "path to file with urls")
	flag.Parse()

	file, err := os.ReadFile(*path)
	if err != nil {
		panic(err.Error())
	}

	links := strings.Split(string(file), "\n")
	respCh := make(chan int)
	errCh := make(chan error)

	for _, raw := range links {
		link := strings.TrimSpace(raw)
		if link == "" {
			continue
		}
		wg.Add(1)
		go ping(link, &wg, respCh, errCh)
	}

	for respCh != nil || errCh != nil {
		select {
		case resp, ok := <-respCh:
			if !ok {
				respCh = nil
				continue
			}
			fmt.Println("OK:", resp)
		case errResp, ok := <-errCh:
			if !ok {
				errCh = nil
				continue
			}
			fmt.Println("ERR:", errResp)
		}
	}
}
