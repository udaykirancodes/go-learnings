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

	for _,link := range links{
		checkLink(link)
	}

}

func checkLink(link string){
	_,err := http.Get(link);
	if err != nil {
		fmt.Println(link," might be down!");
		return 
	}
	fmt.Println(link , " is up!")
}