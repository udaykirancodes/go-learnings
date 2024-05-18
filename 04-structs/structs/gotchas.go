package main

import "fmt"

func main(){

	mySlice := []string{
		"Hi","There","My","Name","Is","Uday!",
	}

	update(mySlice); // this passes by reference. Not pass by value 

	fmt.Println(mySlice); 

}

func update(mySlice []string)  {
	mySlice[0] = "Bye"
}