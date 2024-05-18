# Go Lang

## Hello World Program

```go
// main.go
package main

import "fmt"

func main(){
	fmt.Printf("Hello World");
}
```

## How to Run

```bash
go run main.go
```

## Go CLI Commands

```bash
go build <filename> -> compiles a bunch of go source code files

go run <filename>   -> compiles and execute one or two files

go fmt              -> Formats all the code in each file in the current directory

go install          -> compiles and "Installs" a package

go get              -> download the raw source code of someone else's package

go test             -> Runs any tests associated with the current project

```

## Pakcages

Package is like a Project or WorkSpace

Package is a collection of common source code files

#### Types of Packages

1. Executable - Generates a file that we can run
2. Reusable - Code used as helpers. Good place to put reusable logic

## Variables

```go
var car string = "Ace of spades"

var 	-> variable
car 	-> variable name
string 	-> data type of the variable

```

we have more types

- bool -> true false
- string -> sequence of characters
- int -> Integers
- float64 -> Floating point numbers

## Assigning Varibles

```go
// declaring a variable
var card string = "Ace of Spades";

// another way of declaring
newstring := "ace"

card = "Five of Diamonds"

fmt.Println(newstring)
fmt.Println(card)
```

## Functions in Go

```go
func main(){
	card := newCard();

	fmt.Println(card);
}

func newCard() string{
	return "Five of Diamonds"
}
```

## Arrays & Slices

- Arrays : Fixed length list of things
- Slices : An array that can grow or shrink

```go
func main(){
	// Slice of Cards
	cards := []string{
		"Ace of Diamonds",newCard(),
	}

	// append
	cards = append(cards, "Six of Spades")
	fmt.Println(cards);

	// Printing cards
	for i,card := range cards{
		fmt.Println(i,card);
	}
}
func newCard() string{
	return "Five of Diamonds"
}

```

## Type

```go
type deck []string
* deck is a new type
* deck type extends or borrows the behavious of []string
```

### Receiver Functions

Add Custom Method to the varibles

```go
type deck []string

// this will add a function to the variables which are of type `deck`
func (d deck) print(){
	for i, card := range d {
		fmt.Println(i,card);
	}
}

```

## Return Two seperate values

```go
// function which returns two values
func deal (d deck , handSize int) (deck,deck) {
	return d[:handSize] , d[handSize :]
}

// use it
cards := newDeck()

first , second := deal(cards, 20)

first.print();
fmt.Println()

```

## Write to a file & read from a file

```go
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename , []byte(d.toString()), 0666)
}

func readFromFile(filename string) ([]byte , error) {
	return os.ReadFile(filename);
}
```

## Struct

Data structure. Collection of properties that are related together

```go
package main

import "fmt"

// creating a structure of type person
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
	fmt.Println(udaykiran)
}
```

## Struct with Struct

```go
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
// This is the receiver function
func (p person) print(){
	fmt.Printf("Person : %+v\n",p);
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
	fmt.Printf("Person : %+v\n",uday);
}
```

## Pointers

pointers holds the reference of the variables

- Go is a pass by value language

```go
// This is the receiver function
func (p *person) updateValue(name string){
	(*p).firstnName = name
}
variableName := person {}
varPointer := &variableName
varPointer.updateValue("some value");
/* --- Another Way --- */
// Update FirstName
func (p *person) updateName(name string){
	p.firstname = name
}
```

## Maps

- map is a colletion of key-value pairs
- All the keys have to be the same type
- All the vaue of the maps have to be same type
- key & value can have different types

```go
colors := make(map[string]string)

// declaring with initial values
colours := map[string]string{
	"red":"#ff0000",
	"green":"#random",
}

// insert to map
colours["white"] = "#ffffff";

// delete from map
delete(colours,"white");

fmt.Println(colours,colors,clrs);
```

## Interfaces

Go - If any other type inside our program which has the function called getGreeting() associalted with it that returns a string then, that type is also promoted to be also a type called bot

- I cannot create a value of the Interface whereas i can create values for structs,maps etc

```go
type bot interface {
	getGreeting() string
}
func printGreeting(b bot){
	fmt.Println(b.getGreeting());
}
func main() {
	eb := englishBot{}
	// eb.printGreeting();

	sb := spanishBot{}
	// sp.printGreeting();

	printGreeting(eb)
	printGreeting(sb)
}

```

## Go Routines

A goroutine is a lightweight thread managed by the Go runtime.

## Concurrency

we can have multiple threads executing code. If one thread blocks, another one is picked up and worked on

## Parallelism

Multiple threads executed at the exact same time. Requires multiple CPU's

## Channels

we use channels to make sure that the main routine is aware of each child routines

channel <- 5

- Send the value '5' into this channel

myNumber <- channel

- Wait for a value to be sent into the channel. when we get one, assign the value to 'myNumber'

fmt.Println(<-channel)

- Wait for a value to be sent into the channel. when we get one, log it out immediately

```go
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
```

## Function Literal

Anonymous Functions

```go

func (){

}()

```
