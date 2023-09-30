// Implement a function that, given an array of URLs and a download function,
// downloads all the data from the urls, merges it into a single dictionary of {url:data}
// and return a method with parallelism (concurrency)

package main

import "fmt"

type urlData struct {
	url  string
	data string
}

func main() {
	urls := []string{"url1", "url2", "url3", "url4"}

	c := make(chan urlData)

	for _, url := range urls {
		go get(url, c)
	}

	out := make(map[string]string)

	for i := 0; i < len(urls); i++ {
		kv := <-c
		out[kv.url] = kv.data
	}

	for k, v := range out {
		fmt.Printf("url: %v, data: %v\n", k, v)
	}
}

func get(url string, c chan urlData) {
	data := urlData{
		url:  url,
		data: fmt.Sprintf("data for url %v", url),
	}
	c <- data

}
