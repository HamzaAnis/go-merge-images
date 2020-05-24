package split

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/HamzaAnis/go-merge-images/merge"
	"github.com/HamzaAnis/go-merge-images/models"
	"github.com/HamzaAnis/go-merge-images/utilities"
	"github.com/schollz/progressbar/v3"
)

// Decode image.Image's pixel data into [][]*Pixel
func DecodePixelsArrayFromImage(img image.Image, offsetX, offsetY int, fileName string, parts []models.Part) [][]*models.Pixel {
	log.Println("Decoding", fileName)
	pixels := [][]*models.Pixel{}

	bar := progressbar.Default(int64(img.Bounds().Max.Y * img.Bounds().Max.X))
	// for j := 0; j < len(parts); j++ {
	// 	// pixels = append(pixels)
	// 	pixelsInner := []*models.Pixel{}
	// 	for y := parts[1].Start; y < parts[1].End; y++ {
	// 		for x := 0; x <= img.Bounds().Max.X; x++ {
	// 			p := &models.Pixel{
	// 				Point: image.Point{x + offsetX, y + offsetY},
	// 				Color: img.At(x, y),
	// 			}
	// 			pixelsInner = append(pixelsInner, p)
	// 			bar.Add(1)
	// 		}
	// 	}
	// 	pixels = append(pixels, pixelsInner)
	// }
	count := 0
	count1 := 0
	pixels1 := []*models.Pixel{}
	for y := 0; y <= img.Bounds().Max.Y; y++ {
		for x := 0; x <= img.Bounds().Max.X; x++ {
			p := &models.Pixel{
				Point: image.Point{x + offsetX, y + offsetY},
				Color: img.At(x, y),
			}
			pixels1 = append(pixels1, p)
			bar.Add(1)

		}
		count1++
		if count < len(parts) {
			if count1 == parts[count].YLen {
				println("reset", count)
				println("Length of pixels", len(pixels1))
				pixels = append(pixels, pixels1)
				count1 = 0
				count++
				pixels1 = []*models.Pixel{}
			}
		}
	}
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
		println(part.Start, " ", part.End, " ", part.YLen)
	}
	pixels := DecodePixelsArrayFromImage(img, 0, 0, path, coordinates)
	println("Length is", len(pixels))
	var rectangles []image.Rectangle
	for i := 0; i < len(coordinates); i++ {
		newRect := image.Rectangle{
			Min: img.Bounds().Min,
			Max: image.Point{
				X: img.Bounds().Max.X,
				Y: coordinates[i].YLen,
			},
		}
		rectangles = append(rectangles, newRect)
	}
	finalImages := make([]*image.RGBA, len(rectangles))

	for i, rectangle := range rectangles {
		finalImages[i] = image.NewRGBA(rectangle)
	}
	d, _ := filepath.Split(path)
	for i := 0; i < len(finalImages); i++ {
		bar := progressbar.Default(int64(len(pixels[i])))
		fileName := fmt.Sprintf("%s%d.png", "spilt", i+1)
		splitPath := filepath.Join(d, fileName)
		log.Println("Writing pixels to", splitPath)
		for _, px := range pixels[0] {
			bar.Add(1)
			finalImages[i].Set(
				px.Point.X,
				px.Point.Y,
				px.Color,
			)
		}
		draw.Draw(finalImages[i], finalImages[i].Bounds(), finalImages[i], image.Point{0, 0}, draw.Src)
		out, err := os.Create(splitPath)
		if err != nil {
			return err
		}
		err = png.Encode(out, finalImages[i])
		if err != nil {
			return err
		}
		log.Println("Merge image saved to ", splitPath)
	}
	return nil
}
