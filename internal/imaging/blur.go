package imaging

import (
	"errors"
	"fmt"
	"github.com/mambadev/image-processing/internal/helpers"
	"image"
	"image/color"
	"io"
	"math"
)

func validateAndGatherImage(img image.Image, kernelSize int) ([][]Pixel, error) {
	// the kernel size must be odd to allow the gaussian process to work correctly.
	if kernelSize%2 == 0 {
		return nil, errors.New(fmt.Sprintf("kernelSize cannot be even, kernalSize provided: %d", kernelSize))
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

		return nil, errors.New(msg)
	}

	pixels, _ := GetPixelsFromImage(img)
	return pixels, nil
}

func sumValuesWithinKernel(x, y, kernelSize int, pixels [][]Pixel) Pixel {
	startIdx := x - int(math.Floor(float64(kernelSize)/2)) - 1
	endIdx := startIdx + kernelSize - 1

	startYIdx := y - int(math.Floor(float64(kernelSize)/2)) - 1
	endYIdx := startYIdx + kernelSize - 1

	kernelInnerSize := kernelSize * kernelSize

	result := Pixel{
		R: 0,
		G: 0,
		B: 0,
		A: 0,
	}

	for i := startIdx; i < endIdx; i++ {
		for j := startYIdx; j < endYIdx; j++ {
			result.R += pixels[i][j].R
			result.G += pixels[i][j].G
			result.B += pixels[i][j].B
			result.A += pixels[i][j].A
		}
	}

	return Pixel{
		R: result.R / kernelInnerSize,
		G: result.G / kernelInnerSize,
		B: result.B / kernelInnerSize,
		A: result.A / kernelInnerSize,
	}
}

func BlurMean(file io.Reader, kernelSize int) (image.Image, error) {
	//  our source image which will be used to generate our pixels and
	// apply our blur.
	img, _, _ := image.Decode(file)

	// the target image of the blur, we must have a new image which will be the
	// output of our blur otherwise our math will be broken and not work  as
	// expected.
	targetImg := image.NewRGBA(img.Bounds())
	pixels, _ := validateAndGatherImage(img, kernelSize)

	// iterate over our entire image pixels in blocks of our kernel
	// size. Taking an average of pixel values and applying this
	// back into our new image position.

	// IGNORE EDGES
	// TODO: THIS SKIPS EDGES AND WORKS WITH A PERFECT SIZE FIRST, UPDATE TO INCLUDE EDGES
	startingOffset := int(math.Floor(float64(kernelSize)/2) + 1)

	// TODO: this can be done with go routines to be faster, we are only reading
	// not updating the value and that value will never change allowing us to
	// be very fast.
	for i := startingOffset; i < len(pixels)-startingOffset; i++ {
		for j := startingOffset; j < len(pixels[i])-startingOffset; j++ {
			// now we must sum all the RGB values within our kernel size
			// and then go and divide this by our total kernel size which
			// would be kernelSize x kernelSize. This would be our new
			// RGB values for the center pixel.
			newPixel := sumValuesWithinKernel(i, j, kernelSize, pixels)

			targetImg.SetRGBA(j, i, color.RGBA{
				R: uint8(newPixel.R),
				G: uint8(newPixel.G),
				B: uint8(newPixel.B),
				A: uint8(newPixel.A),
			})

			// fmt.Printf("center pixel %v, size: %d\n", pixel, kernelInnerSize)
		}
	}

	return targetImg, nil
}

func BlurGaussian(file io.Reader, kernelSize int) error {
	//  our source image which will be used to generate our pixels and
	// apply our blur.
	img, _, _ := image.Decode(file)

	// the target image of the blur, we must have a new image which will be the
	// output of our blur otherwise our math will be broken and not work  as
	// expected.
	//	targetImg := image.NewRGBA(img.Bounds())

	pixels, _ := validateAndGatherImage(img, kernelSize)

	fmt.Println(pixels)
	return nil
}
