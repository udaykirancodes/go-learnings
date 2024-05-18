package main

import "fmt"

type person struct{
	firstname string 
	lastname string 
}

func main(){

	// creating a struct 
	uday := person{"uday","kiran"}
	// anther way 
	udaykiran := person{
		firstname : "uday",
		lastname: "kiran",
	}
	fmt.Println(uday)
	fmt.Printf("User Struct : %+v\n",udaykiran)


	var alex person;
	alex.firstname =
	 "udaykiran" 
	fmt.Println(alex.firstname)
}