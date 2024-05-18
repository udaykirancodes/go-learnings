package main

import (
	"fmt"
)

func main()  {

	// declare map <string,string>
	var clrs map[string]string

	colors := make(map[string]string)


	// declaring with initial values 
	colours := map[string]string{
		"red":"#ff0000",
		"green":"#random",
	}	
	fmt.Println("Printing : ");
	// Iterate in the map 
	for key,value := range colours {
		fmt.Println(key+" "+value);
	}

	// insert to map 
	colours["white"] = "#ffffff";
	
	// delete from map 
	delete(colours,"white");


	fmt.Println(colours,colors,clrs);
}