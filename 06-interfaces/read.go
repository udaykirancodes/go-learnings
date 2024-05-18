package main

import (
	"io"
	"os"
)

func main() {
	arr :=  os.Args;

	if(len(arr) < 2) {
		return
	}

	file,err := os.Open("text.txt")
	if err != nil {
		return
	}

	io.Copy(os.Stdout , file)

}