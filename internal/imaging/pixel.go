package imaging

import (
	"image"
)

// Pixel struct example
type Pixel struct {
	R int
	G int
	B int
	A int
}

func RgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 255), int(g / 255), int(b / 255), int(a / 255)}
}

// GetPixelsFromImage will take in a file reader and decode the content into an
// imaging. The imaging will be iterated and converted into a Pixel array which
//will contain the RGBA values of the given pixel at that point.
func GetPixelsFromImage(img image.Image) ([][]Pixel, error) {
	// pull the bounds of the imaging to determine the range of which
	// it will be iterating to gather the pixels. Bounds returns
	// our max width and height.
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel

	for y := 0; y < height; y++ {
		var row []Pixel

		for x := 0; x < width; x++ {
			row = append(row, RgbaToPixel(img.At(x, y).RGBA()))
		}

		pixels = append(pixels, row)
	}

	return pixels, nil
}
