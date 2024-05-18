package main

import (
	"fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	fmt.Print(len(d))
	if len(d) != 16 {
		t.Errorf("Expected deck length of 40, but got %v", len(d))
	}
}