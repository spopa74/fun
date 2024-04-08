package datastructures

import "fmt"

func TestBitwise() {

	fmt.Printf("10 in binary: %b\n", 10)
	fmt.Printf("first bit of 10: %b\n", 10&(1))
	fmt.Printf("second bit of 10: %b\n", 10&(1<<1)>>1)
	fmt.Printf("third bit of 10: %b\n", 10&(1<<2)>>2)
	fmt.Printf("fourth bit of 10: %b\n", 10&(1<<3)>>3)
	// fmt.Printf("%b", 10)
	// fmt.Printf("%b", 10)
	// fmt.Printf("%b", 10)

}
