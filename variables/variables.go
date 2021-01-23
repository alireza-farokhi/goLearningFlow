package main

import (
	"fmt"
)

// can be used etire file
var h = "use var keyword to declare a variable outside of a function"

/***
	if uncommnet above line we'll get an errorl
***/
// h := "we'll get error, becuase we can't use shorthand declaration outside of a function body"

func main() {
	// shorthand variable (identifire) declaration and assignment
	// should use just inside a function body
	// inside declared variables must used somwhere, otherwise compiler shows error
	// the scope of this variable is in just this very function
	x := 42
	fmt.Println(x)

	x = 99
	fmt.Println(x)

	var y = "olemon403"
	fmt.Println(y)

	// outsdie declared variables can be unused
	fmt.Println(h)

	foo()
}

func foo() {
	fmt.Println("again:", h)
}
