package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(double(100000000000000000))

	first, _ := parseName("Dave Nelson")
	fmt.Println(first)

	greet := func(name string) {
		fmt.Println("Hello", name)
	}

	greet(first)
}

func double(n int) int {
	return n + n
}

func parseName(name string) (first, last string) {
	parsed := strings.Split(name, " ")
	return parsed[0], parsed[1]
}
