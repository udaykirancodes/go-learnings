package main

import "fmt"

// contact info struct
type contactInfo struct {
	email string
	zipCode string
}

// person Info struct 
type person struct {
	firstname string
	lastname string 
	contact contactInfo
}

// Receiver Function 
func (p person) print(){
	fmt.Printf("Person : %+v\n",p);
}
// Update FirstName 
func (p *person) updateName(name string){
	p.firstname = name 
}

func main(){

	uday := person{
		firstname: "udaykiran",
		lastname : "bandarugalla",
		contact : contactInfo{
			email : "officialudaykiran@gmail.com",
			zipCode: "502248",
		},
	}
	uday.updateName("hello");
	uday.print();
}