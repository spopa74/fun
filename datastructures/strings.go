package datastructures

import "fmt"

func TestStrings() {
	someString := "meh"
	fmt.Println(someString)

	var b byte = someString[2]
	fmt.Println(b)

	var h string = string(b)
	fmt.Println(h)

	var conc string = "me" + h + "2"
	fmt.Println(conc)

}
