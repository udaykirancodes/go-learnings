package main

import "fmt"


func main(){


	fmt.Println("Hello");


	func(){
		fmt.Println("World");
	}()

}