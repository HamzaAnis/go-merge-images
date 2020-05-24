package split

import (
	"log"

	"github.com/HamzaAnis/go-merge-images/merge"
)

func SplitImage(path string) error {
	img, _, err := merge.OpenAndDecode(path)
	if err != nil {
		return err
	}
	// pixels := merge.DecodePixelsFromImage(img, 0, 0, path)
	log.Println("It is ")
	log.Println(img.Bounds().Max.Y)
	return nil
}
