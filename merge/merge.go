package merge

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/HamzaAnis/go-merge-images/models"
	"github.com/HamzaAnis/go-merge-images/utilities"
	"github.com/schollz/progressbar/v3"
)

// Keep it DRY so don't have to repeat opening file and decode
func OpenAndDecode(filepath string) (image.Image, string, error) {
	imgFile, err := os.Open(filepath)
	if err != nil {
		return nil, "", err
	}
	defer imgFile.Close()
	img, format, err := image.Decode(imgFile)
	if err != nil {
		return nil, "", err
	}
	return img, format, nil
}

// Decode image.Image's pixel data into []*Pixel
func DecodePixelsFromImage(img image.Image, offsetX, offsetY int, fileName string) []*models.Pixel {
	log.Println("Decoding", fileName)
	pixels := []*models.Pixel{}
	bar := progressbar.Default(int64(img.Bounds().Max.Y * img.Bounds().Max.X))
	for y := 0; y <= img.Bounds().Max.Y; y++ {
		for x := 0; x <= img.Bounds().Max.X; x++ {
			p := &models.Pixel{
				Point: image.Point{x + offsetX, y + offsetY},
				Color: img.At(x, y),
			}
			pixels = append(pixels, p)
			bar.Add(1)
		}
	}
	return pixels
}

func MergeImages(directory models.Directory) error {
	log.Println("Processing", directory.DirectoryPath)

	filesD := directory.Files
	var files []string
	files = append(files, filesD[len(filesD)-1])
	for i := 0; i < len(filesD)-1; i++ {
		files = append(files, filesD[i])
	}
	log.Println(" All the images found in the sorted order [", directory.DirectoryPath, "]")
	for _, a := range files {
		_, f := filepath.Split(a)
		log.Println(filepath.Split(f))
	}
	images := []image.Image{}

	for _, file := range files {
		img, _, err := OpenAndDecode(file)
		if err != nil {
			return err
		}
		images = append(images, img)
	}
	pixels := []*models.Pixel{}
	_, f := filepath.Split(files[0])
	var temp = DecodePixelsFromImage(images[0], 0, 0, f)
	pixels = append(pixels, temp...)
	maxY := images[0].Bounds().Max.Y
	for i := 1; i < len(images); i++ {
		offset := 0
		for j := 0; j < i; j++ {
			offset += images[j].Bounds().Max.Y
		}
		_, f := filepath.Split(files[i])
		var temp1 = DecodePixelsFromImage(images[i], 0, offset, f)
		pixels = append(pixels, temp1...)
		maxY += images[i].Bounds().Max.Y
	}
	pixelSum := []*models.Pixel{}

	for _, pixel := range pixels {
		pixelSum = append(pixelSum, pixel)
	}
	mergePath := filepath.Join(directory.DirectoryPath, "merged.png")
	log.Println("Writing pixels to merge.png")
	bar := progressbar.Default(int64(len(pixelSum)))

	newRect := image.Rectangle{
		Min: images[0].Bounds().Min,
		Max: image.Point{
			X: images[1].Bounds().Max.X,
			Y: maxY,
		},
	}
	finImage := image.NewRGBA(newRect)
	for _, px := range pixelSum {
		bar.Add(1)
		finImage.Set(
			px.Point.X,
			px.Point.Y,
			px.Color,
		)
	}
	draw.Draw(finImage, finImage.Bounds(), finImage, image.Point{0, 0}, draw.Src)
	out, err := os.Create(mergePath)
	if err != nil {
		return err
	}
	err = png.Encode(out, finImage)
	if err != nil {
		return err
	}
	log.Println("Merge image saved to ", mergePath)
	for _, file := range files {
		utilities.DeleteFile(file)
		if err != nil {
			return err
		}
	}
	return nil
}
