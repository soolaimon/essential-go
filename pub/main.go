package main

import (
	"fmt"
	"github.com/soolaimon/essential-go/pub/pages"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Not enough arguments")
	}
	filename := os.Args[1]
	fmt.Println(filename)
	p, err := pages.NewPage(filename)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(p)

}
