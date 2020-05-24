package cmd

import (
	"log"

	"github.com/HamzaAnis/go-merge-images/merge"
	"github.com/HamzaAnis/go-merge-images/split"
	"github.com/HamzaAnis/go-merge-images/utilities"
)

func Run() {
	header := `
	██╗███╗   ███╗ █████╗  ██████╗ ███████╗    ███╗   ███╗███████╗██████╗  ██████╗ ███████╗
	██║████╗ ████║██╔══██╗██╔════╝ ██╔════╝    ████╗ ████║██╔════╝██╔══██╗██╔════╝ ██╔════╝
	██║██╔████╔██║███████║██║  ███╗█████╗      ██╔████╔██║█████╗  ██████╔╝██║  ███╗█████╗  
	██║██║╚██╔╝██║██╔══██║██║   ██║██╔══╝      ██║╚██╔╝██║██╔══╝  ██╔══██╗██║   ██║██╔══╝  
	██║██║ ╚═╝ ██║██║  ██║╚██████╔╝███████╗    ██║ ╚═╝ ██║███████╗██║  ██║╚██████╔╝███████╗
	╚═╝╚═╝     ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚══════╝    ╚═╝     ╚═╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚══════╝
	
	Developed by:  Hamza Anis																					                                                                                       
																					   `
	println(header)
	paths, err := utilities.GetDirectories("/Users/macbookpro/Desktop/Programming/test")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("All the subdirectories found")
	for _, elem := range paths {
		log.Println(elem)
	}
	if utilities.AskForConfirmation("Do you want to start the merging (Y/N)?") {
		directories, err := utilities.GetProcessedDirectories("/Users/macbookpro/Desktop/Programming/test")
		if err != nil {
			log.Fatal(err)
		}

		for _, directory := range directories {
			err = merge.MergeImages(directory)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if utilities.AskForConfirmation("Do you want to start the splitting (Y/N)?") {
		directories, err := utilities.GetProcessedDirectories("/Users/macbookpro/Desktop/Programming/test")
		if err != nil {
			log.Fatal(err)
		}

		for _, directory := range directories {
			err = merge.MergeImages(directory)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	err = split.SplitImage("/Users/macbookpro/Desktop/Programming/go-merge-images/test/test1/merged.png")
	if err != nil {
		log.Fatal(err)
	}
}
