package models

import (
	"image"
	"image/color"
)

type Directory struct {
	DirectoryPath string
	Files         []string
}

// Create a struct to deal with pixel
type Pixel struct {
	Point image.Point
	Color color.Color
}

type Part struct {
	Start int
	End   int
	YLen  int
}
