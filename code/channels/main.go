package main

import(
	"fmt"
	"time"
	"net/http"
)
func main(){
	links :=[]string{
		"http://twitter.com",
		"http://facebook.com",
		"http://netflix.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://amazon.com",
	}
	
	c := make(chan string)

	for _,link := range links {
		go checkLink(link,c)
	}
	
	for l:= range c {
		go func() {
			time.Sleep(5 *time.Second)
			checkLink(l,c)
		}()
	}
}


func checkLink(link string,c chan string) {
	_, err := http.Get(link)
	if err !=nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
