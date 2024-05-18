package main

import "fmt"

type bot interface {
	getGreeting() string 
}

type englishBot struct {}

type spanishBot struct {}


func main() {
	fmt.Println("Interfaces");

	eb := englishBot{}
	// eb.printGreeting();
	
	sb := spanishBot{}
	// sp.printGreeting();

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting());
}

func (eb englishBot) getGreeting() string {
	return "Hello There!"
}
func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
func (eb englishBot) printGreeting()  {
	fmt.Println(eb.getGreeting());
}
func (sb spanishBot) printGreeting()  {
	fmt.Println(sb.getGreeting())
}