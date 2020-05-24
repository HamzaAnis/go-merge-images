package split

import (
	"image"
	"log"

	"github.com/HamzaAnis/go-merge-images/merge"
	"github.com/HamzaAnis/go-merge-images/models"
	"github.com/HamzaAnis/go-merge-images/utilities"
)

// Decode image.Image's pixel data into []*Pixel
func DecodePixelsArrayFromImage(img image.Image, offsetX, offsetY int, fileName string, parts []int) [][]*models.Pixel {
	log.Println("Decoding", fileName)

	pixels := [][]*models.Pixel{}
	// bar := progressbar.Default(int64(img.Bounds().Max.Y * img.Bounds().Max.X))
	for i := 0; i < len(parts); i++ {
		for j := 0; j < parts[i]; j++ {

		}
	}
	// pixels := []*models.Pixel{}
	// bar := progressbar.Default(int64(img.Bounds().Max.Y * img.Bounds().Max.X))
	// for y := 0; y <= img.Bounds().Max.Y; y++ {
	// 	for x := 0; x <= img.Bounds().Max.X; x++ {
	// 		p := &models.Pixel{
	// 			Point: image.Point{x + offsetX, y + offsetY},
	// 			Color: img.At(x, y),
	// 		}
	// 		pixels = append(pixels, p)
	// 		bar.Add(1)
	// 	}
	// }
	return pixels
}
func SplitImage(path string) error {
	img, _, err := merge.OpenAndDecode(path)
	if err != nil {
		return err
	}
	coordinates, err := utilities.SplitParts(img.Bounds().Max.Y, 16300)
	if err != nil {
		return err
	}
	for _, part := range coordinates {
		println(part.Start, " ", part.End)
	}
	// pixels := merge.DecodePixelsFromImage(img, 0, 0, path)
	// println("Length is", len(pixels))
	// var rectangles []image.Rectangle
	// for i := 0; i < len(coordinates); i++ {
	// 	newRect := image.Rectangle{
	// 		Min: img.Bounds().Min,
	// 		Max: image.Point{
	// 			X: img.Bounds().Max.X,
	// 			Y: 16300,
	// 		},
	// 	}
	// 	rectangles = append(rectangles, newRect)
	// }

	return nil
}
