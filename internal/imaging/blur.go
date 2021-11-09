package imaging

import (
	"errors"
	"fmt"
	"github.com/mambadev/image-processing/internal/helpers"
	"image"
	"io"
)

func BlurGaussian(file io.Reader, kernelSize int) error {
	// the kernel size must be odd to allow the gaussian process to work correctly.
	if kernelSize%2 == 0 {
		return errors.New(fmt.Sprintf("kernelSize cannot be even, kernalSize provided: %d", kernelSize))
	}

	img, _, err := image.Decode(file)

	if err != nil {
		return err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	maxKernelSize := helpers.Min(width, height)

	// don't allow the kernel size ot be greater than the image size, (height or width),
	// which ever is smaller, since we must allow the blur process to function within
	// the image range otherwise it will always have flipped values.
	if maxKernelSize < kernelSize {
		msg := fmt.Sprintf("kernelSize cannot be greater than the smallest width "+
			"or height value, kernalSize provided: %d, max: %d", kernelSize, maxKernelSize)

		return errors.New(msg)
	}

	pixels, _ := GetPixelsFromImage(img)
	fmt.Println(pixels)

	return nil
}
