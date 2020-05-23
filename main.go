package main

import (
	"log"

	"github.com/HamzaAnis/go-merge-images/utilities"
)

func main() {
	// paths, err := utilities.GetDirectories("/Users/macbookpro/Desktop/Programming/test")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, elem := range paths {
	// 	println(elem)
	// }

	directories, err := utilities.GetProcessedDirectories("/Users/macbookpro/Desktop/Programming/test")
	if err != nil {
		log.Fatal(err)
	}

	for _, directory := range directories {
		log.Println("For ", directory.DirectoryPath)
		for _, file := range directory.Files {
			log.Println(file)
		}
		print()
	}
}
