package main

import (
	"fmt"
	"io"
	"net/http"
)

// When passing a second parameter to the make function when creating a channel we can specify the buffer size
// c := make(chan int, 1) - This creates a buffered channel with a capacity of 1.
// Normally, channels are synchronous; both sides of the channel will wait until the other side is ready.
// A buffered channel is asynchronous; sending or receiving a message will not wait unless the channel
// is already full. If the channel is full, then sending will wait until there is room for at least one more int

type HomePageSize struct {
	URL  string
	Size int
	Err  error
}

func main() {
	urls := []string{
		"https://www.apple.com",
		"https://www.amazon.com",
		"https://www.google.com",
		"https://www.microsoft.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.linkedin.com",
		"https://www.twitter.com",
		"https://www.tiktok.com",
		"https://www.youtube.com",
	}

	results := make(chan HomePageSize)
	for _, url := range urls {
		// unnamed function that is immediately invoked.
		// This is a common pattern with goroutines, but also notice that we defined the function as taking a
		// single parameter (the url). The reason this function doesn’t reference the url directly
		// (which it is allowed to do) is that the rules of closure are such that if we did that, all
		// four goroutines would probably end up seeing the same value for url. This is because
		// url is changed by the for loop.
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				results <- HomePageSize{
					URL:  url,
					Err:  err,
					Size: 0,
				}
				return
			}
			defer res.Body.Close()

			bs, err := io.ReadAll(res.Body)
			if err != nil {
				results <- HomePageSize{
					URL:  url,
					Err:  err,
					Size: 0,
				}
				return
			}
			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
				Err:  nil,
			}
		}(url)
	}

	var biggest HomePageSize
	errors := make([]HomePageSize, 0)
	// technic to iterate over a slice without the traditionally 'for loop'
	for range urls {
		result := <-results
		if result.Err == nil && result.Size > biggest.Size {
			biggest = result
		}
		if result.Err != nil {
			errors = append(errors, result)
		}
	}

	fmt.Printf("The biggest home page: %s with size %d\n", biggest.URL, biggest.Size)
	fmt.Println("Home pages with errors:")
	for _, err := range errors {
		fmt.Printf("\t - URL: %s, Error: %s\n", err.URL, err.Err)
	}
}
