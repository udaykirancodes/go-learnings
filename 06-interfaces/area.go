package main

import "fmt"

type shape interface{
	getArea() float64
}

type square struct {
	sideLength float64 
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength;
}


type triangle struct {
	base float64
	height float64
}

func (t triangle) getArea() float64 {
	return 0.5 * (t.base) * (t.height)
}

func printArea(s shape){
	fmt.Println(s.getArea());
}

func main(){
	t := triangle{
		base: 12,
		height: 12,
	}
	printArea(t);
	s := square{
		sideLength: 12,
	}
	printArea(s);
}