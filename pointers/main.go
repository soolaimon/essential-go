package main

import (
	"fmt"
)

func main() {
	num := 5
	double(num)
	fmt.Println(num)
	triple(&num)
	fmt.Println(num)

	x, y := 20, 40
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)

}

func double(n int) {
	n = n * 2
}

func triple(n *int) {
	*n = *n * 3
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}
