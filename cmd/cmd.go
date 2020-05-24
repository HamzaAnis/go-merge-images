package cmd

import (
	"log"

	"github.com/HamzaAnis/go-merge-images/merge"
	"github.com/HamzaAnis/go-merge-images/utilities"
)

func Run() {
	paths, err := utilities.GetDirectories("/Users/macbookpro/Desktop/Programming/test")
	if err != nil {
		log.Fatal(err)
	}
	for _, elem := range paths {
		println(elem)
	}

	directories, err := utilities.GetProcessedDirectories("/Users/macbookpro/Desktop/Programming/test")
	if err != nil {
		log.Fatal(err)
	}

	// for _, directory := range directories {
	// 	log.Println("Processing ", directory.DirectoryPath)
	// 	err := merge.MergeImages(directory)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	err = merge.MergeImages(directories[0])
	if err != nil {
		log.Fatal(err)
	}
}
