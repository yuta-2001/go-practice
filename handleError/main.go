package main

import (
	"net/http"
	"fmt"
)

type Result struct {
	Error error
	Response *http.Response
}


func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
	
			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}
	
	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://www.yahoo.com"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Println("Error: ", result.Error)
			continue
		}
		fmt.Println("Response: ", result.Response.Status)
	}
}
