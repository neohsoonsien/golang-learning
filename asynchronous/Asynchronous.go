
package asynchronous

import (
	"fmt"
	"net/http"
	"time"
)

func Asynchronous() {
	// A slice of example websites
	urls := []string{
		"https://www.twitter.com/",
		"https://www.google.com",
		"https://www.facebook.com/",
		"https://www.youtube.com/",
	}
	for _, url := range urls {
		go checkUrl(url)
	}
	// need to wait for all the asynchronous process to complete
	// else the main function will exit immediately
	time.Sleep(3 * time.Second) 
}

//checks and prints a message if a website is up or down
func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down !!!")
		return
	}
	fmt.Println(url, "is up and running.")
}