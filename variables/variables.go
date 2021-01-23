package main

import (
	"fmt"
)

var a string = "42"
var b string = "some text"
var c string
var d = "some another"
var e = `text with backqoute`
var f string = `take advatage of "back tick" to us double qoutes inside of a string`
var g string = ""
var h string = ``
var i = ""
var j = ``

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println("\n")

	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println("\n")

	fmt.Println(c)
	fmt.Printf("%T\n", c)
	fmt.Println("\n")

	fmt.Println(d)
	fmt.Printf("%T\n", d)
	fmt.Println("\n")

	fmt.Println(e)
	fmt.Printf("%T\n", e)
	fmt.Println("\n")

	fmt.Println(f)
	fmt.Printf("%T\n", f)
	fmt.Println("\n")

	fmt.Println(g)
	fmt.Printf("%T\n", g)
	fmt.Println("\n")

	fmt.Println(h)
	fmt.Printf("%T\n", h)
	fmt.Println("\n")

	fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Println("\n")

	fmt.Println(j)
	fmt.Printf("%T\n", j)
	fmt.Println("\n")
}
