package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	ch := make(chan string)
	links := []string{"http://www.facebook.com", "http://www.google.com"}

	go checkLink(links, ch)

	for i := range ch {
		fmt.Println("recived -> ", i)
	}

}

func checkLink(l []string, ch chan string) {
	var wg sync.WaitGroup
	for _, siteurl := range l {
		trigger := func() {
			_, err := http.Get(siteurl)
			if err != nil {
				ch <- "error"
			}
			fmt.Println(siteurl, " is up ")
			ch <- siteurl + " is up"
			wg.Done()
		}
		wg.Add(1)
		go trigger()
	}
	wg.Wait()
	close(ch)
	fmt.Println("Completed!!!")
}
