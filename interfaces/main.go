package main

import (
	"fmt"
)

type Animal interface {
	Pet()
	Name() string
}

func Compliment(a Animal) {
	fmt.Println("Great Job", a.Name())
	a.Pet()
}

type Cat struct {
	name string
}

func (c Cat) Name() string {
	return c.name
}

func (c Cat) Pet() {
	fmt.Println("prrrr")
}

type Dog struct {
	name string
}

func (d Dog) Pet() {
	fmt.Println("woof woof")
}

func (d Dog) Name() string {
	return d.name
}

type Liger struct {
	name string
}

func (l Liger) Name() string {
	return l.name
}

func (l Liger) Pet() {
	fmt.Println("growl")
}

func main() {

	c := Cat{"johnny larry"}
	Compliment(c)

	d := Dog{"boomer nelson"}
	Compliment(d)

	l := Liger{"max von whatever"}
	Compliment(l)
}
