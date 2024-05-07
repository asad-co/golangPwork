package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Process start:", start)

	links := []string{
		"https://go.dev/learn/",
		"https://www.netflix.com",
		"https://go.dev/learn/",
		"https://www.netflix.com",
		"https://go.dev/learn/",
		"https://www.netflix.com",
		"https://go.dev/learn/",
		"https://www.netflix.com",
		"https://www.twitter.com",
		"https://www.xhamster.com",
	}
	checkUrls(links)
	fmt.Printf("Completed the code process in : %f seconds", time.Since(start).Seconds())

}

func checkUrls(urls []string) {
	c := make(chan string)
	var wg sync.WaitGroup
	for _, link := range urls {
		wg.Add(1)
		go checkUrl(link, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()
	for msg := range c {
		fmt.Println(msg)
	}
}
func checkUrl(url string, c chan string, wg *sync.WaitGroup) {
	defer (*wg).Done()
	_, err := http.Get(url)

	if err != nil {
		c <- "We could not reach: " + url
		return
	}
	c <- "Successfully reached the website" + url
}
