package imaging

import (
	"errors"
	"fmt"
	"github.com/stephensli/image-processing/internal/helpers"
	"image"
	"image/color"
)

type Blur interface {
	Blur() (image.Image, error)
	getUpdatedPixel(x, y int, pixels [][]Pixel) color.RGBA
	validateAndGetImagePixels() ([][]Pixel, error)
}

type BlurAction struct {
	KernelSize int
	Iterations int
	Image      image.Image
}

func (b *BlurAction) validateAndGetImagePixels() ([][]Pixel, error) {
	// the kernel size must be odd to allow the gaussian process to work correctly.
	if b.KernelSize%2 == 0 {
		return nil, errors.New(fmt.Sprintf("kernelSize cannot be even, kernalSize provided: %d", b.KernelSize))
	}

	bounds := b.Image.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	maxKernelSize := helpers.Min(width, height)

	// don't allow the kernel size ot be greater than the image size, (height or width),
	// which ever is smaller, since we must allow the blur process to function within
	// the image range otherwise it will always have flipped values.
	if maxKernelSize < b.KernelSize {
		msg := fmt.Sprintf("kernelSize cannot be greater than the smallest width "+
			"or height value, kernalSize provided: %d, max: %d", b.KernelSize, maxKernelSize)

		return nil, errors.New(msg)
	}

	pixels, _ := GetPixelsFromImage(b.Image)
	return pixels, nil
}
