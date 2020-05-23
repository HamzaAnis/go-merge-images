package main

import (
	"log"

	"github.com/HamzaAnis/go-merge-images/utilities"
)

func main() {
	paths, err := utilities.GetDirectories("/Users/macbookpro/Desktop/Programming/test")
	if err != nil {
		log.Fatal(err)
	}
	for _, elem := range paths {
		println(elem)
	}
	filepaths, err := utilities.GetFiles("/Users/macbookpro/Desktop/Programming/test/test1")
	if err != nil {
		log.Fatal(err)
	}
	for _, elem := range filepaths {
		println(elem)
	}
}
