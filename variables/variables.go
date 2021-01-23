package main

import (
	"fmt"
)

var y = 42

func main() {
	fmt.Println(y)
	// get type of a certain variable using printf identifier of fmt package
	fmt.Printf("%T\n", y)
}
