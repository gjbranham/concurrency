package main

import "fmt"

// Implement a function that, given an array of URLs and a download function,
// downloads all the data from the urls, merges it into a single dictionary of {url:data}
// and return a method with parallelism (concurrency)

type urlData struct {
	url  string
	data string
}

func main() {
	urls := []string{"url1", "url2", "url3", "url4"}

	c := make(chan urlData)

	go get(urls[:len(urls)/2], c)
	go get(urls[len(urls)/2:], c)

	out := make(map[string]string)

	for i := 0; i < len(urls); i++ {
		kv := <-c
		out[kv.url] = kv.data
	}

	for k, v := range out {
		fmt.Printf("url: %v, data: %v\n", k, v)
	}

}

func get(urls []string, c chan urlData) {
	for _, u := range urls {
		data := urlData{
			url:  u,
			data: fmt.Sprintf("data for url %v", u),
		}
		c <- data
	}
}
