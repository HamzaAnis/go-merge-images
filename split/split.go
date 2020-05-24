package split

import (
	"image"

	"github.com/HamzaAnis/go-merge-images/merge"
	"github.com/HamzaAnis/go-merge-images/utilities"
)

func SplitImage(path string) error {
	img, _, err := merge.OpenAndDecode(path)
	if err != nil {
		return err
	}
	// pixels := merge.DecodePixelsFromImage(img, 0, 0, path)
	coordinates, err := utilities.SplitParts(img.Bounds().Max.Y, 16300)
	if err != nil {
		return err
	}
	pixels := merge.DecodePixelsFromImage(img, 0, 0, path)
	println("Length is", len(pixels))
	var rectangles []image.Rectangle
	for i := 0; i < len(coordinates); i++ {
		newRect := image.Rectangle{
			Min: img.Bounds().Min,
			Max: image.Point{
				X: img.Bounds().Max.X,
				Y: 16300,
			},
		}
		rectangles = append(rectangles, newRect)
	}

	return nil
}
