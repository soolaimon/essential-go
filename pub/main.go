package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Not enough arguments")
	}
	filename := os.Args[1]
	fmt.Println(filename)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))

	data = blackfriday.MarkdownCommon(data)
	fmt.Println(string(data))
}
