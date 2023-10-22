package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://golang.org",
		"https://stackoverflow.com",
		"https://google.com",
		"https://amazon.com",
		"https://facebook.com",
	}

	c := make(chan string)

	for _, url := range links {
		go checkLink(url, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	//  infitiny loop
	for l := range c {
		go func(urlFromChan string) {
			time.Sleep(5 * time.Second)
			checkLink(urlFromChan, c)
		}(l)
	}

}

func checkLink(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, " ----> [ might be down ]")
		// log.Fatal(err)
		c <- url
		return
	}

	fmt.Println(url, "-----> [ is up ]")
	c <- url
}
