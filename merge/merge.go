package merge

import (
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/prometheus/common/model"
)

// Keep it DRY so don't have to repeat opening file and decode
func OpenAndDecode(filepath string) (image.Image, string, error) {
	imgFile, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer imgFile.Close()
	img, format, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}
	return img, format, nil
}

// Decode image.Image's pixel data into []*Pixel
func DecodePixelsFromImage(img image.Image, offsetX, offsetY int) []*Pixel {
	pixels := []*Pixel{}
	for y := 0; y <= img.Bounds().Max.Y; y++ {
		for x := 0; x <= img.Bounds().Max.X; x++ {
			p := &Pixel{
				Point: image.Point{x + offsetX, y + offsetY},
				Color: img.At(x, y),
			}
			pixels = append(pixels, p)
		}
	}
	return pixels
}

func MergeImages(directory model.Directory) err {
	files := []string{}
	files = append(files, "screencapture-comic-naver-webtoon-detail-nhn-2020-05-22-20_16_04.png")
	files = append(files, "screencapture-comic-naver-webtoon-detail-nhn-2020-05-22-20_16_04-2.png")
	files = append(files, "screencapture-comic-naver-webtoon-detail-nhn-2020-05-22-20_16_04-3.png")
	files = append(files, "screencapture-comic-naver-webtoon-detail-nhn-2020-05-22-20_16_04-4.png")
	files = append(files, "screencapture-comic-naver-webtoon-detail-nhn-2020-05-22-20_16_04-5.png")
	for _, a := range files {
		println(a)
	}
	images := []image.Image{}

	for _, file := range files {
		img, _, err := OpenAndDecode(file)
		if err != nil {
			panic(err)
		}
		images = append(images, img)
	}
	pixels := []*Pixel{}
	var temp = DecodePixelsFromImage(images[0], 0, 0)
	pixels = append(pixels, temp...)
	maxY := images[0].Bounds().Max.Y
	for i := 1; i < len(images); i++ {
		offset := 0
		for j := 0; j < i; j++ {
			offset += images[j].Bounds().Max.Y
		}
		var temp1 = DecodePixelsFromImage(images[i], 0, offset)
		pixels = append(pixels, temp1...)
		maxY += images[i].Bounds().Max.Y
	}
	println("Max Y ", maxY)
	pixelSum := []*Pixel{}

	for _, pixel := range pixels {
		pixelSum = append(pixelSum, pixel)
	}
	newRect := image.Rectangle{
		Min: images[0].Bounds().Min,
		Max: image.Point{
			X: images[1].Bounds().Max.X,
			Y: maxY,
		},
	}
	finImage := image.NewRGBA(newRect)
	for _, px := range pixelSum {
		finImage.Set(
			px.Point.X,
			px.Point.Y,
			px.Color,
		)
	}
	draw.Draw(finImage, finImage.Bounds(), finImage, image.Point{0, 0}, draw.Src)
	out, err := os.Create("./merged.png")
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	err = png.Encode(out, finImage)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}
