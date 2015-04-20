package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y int
}

func NewPoint(x, y int) point {
	return point{x, y}
}

type rect struct {
	pos    point
	width  int
	height int
}

type circle struct {
	radius float64
}

func (c circle) circumference() float64 {
	return math.Pi * c.radius * 2

}

func (r rect) area() int {
	return r.width * r.height
}

func main() {

	p := point{20, 40}
	fmt.Println(p)

	r := rect{
		pos:    NewPoint(20, 30),
		width:  200,
		height: 400,
	}
	fmt.Println(r)
	fmt.Println(r.area())
	c := circle{
		radius: 20.0,
	}

	fmt.Println(c)
	fmt.Println(c.circumference())

}
