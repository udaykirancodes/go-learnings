package main

import (
	"fmt"
	"net/http"
)

func main(){

	links := []string {
		"https://google.com",
		"https://instagram.com",
		"https://facebook.com",
	}

	c := make(chan string);

	for _,link := range links{
		go checkLink(link,c)
	}
	
	for i:=0 ; i<len(links) ; i++ {
		fmt.Println(<-c)
	}

}

func checkLink(link string,c chan string){ 
	_,err := http.Get(link);
	if err != nil {
		fmt.Println(link," might be down!");
		c <- "Might be down I think"
		return 
	}
	fmt.Println(link , " is up!")
	c <- "Yup It's Up!"
}