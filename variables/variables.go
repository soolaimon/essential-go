package main

import "fmt"

var (
	author  = "@codegangsta"
	Version = "0.0.1"
)

const (
	CCVisa       = "Visa"
	CCMastercard = "MasterCard"
	CCAmex       = "American Express"
)

func main() {
	var greeting string = "Hellow World!"
	var a, b, c int = 1, 2, 3
	fmt.Println(greeting, a, b, c)

	var name = "Dave Nelson"
	var d, e, f = 1, 2.0, "Three"
	fmt.Println(name, d, e, f)

	course := "Essential Go"
	x, y, z := 1, 2, 3
	fmt.Println(course, x, y, z)

	fmt.Fprintf("There are", y, " ", CCVisa, "s here. This is", true)
}
