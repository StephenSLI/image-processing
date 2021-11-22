package imaging

import (
	"image"
	"image/color"
	"math"
	"sync"
)

type BlurActionMean struct {
	BlurAction
}

// getUpdatedPixel will iterate within the given kernel value range and
// determine the total mean value within the selected kernel size. The returned value being
// the new pixel which could be placed within the center of the kernel of the new image
func (b *BlurActionMean) getUpdatedPixel(x, y int, pixels [][]Pixel) color.RGBA {
	startIdx := x - int(math.Floor(float64(b.KernelSize)/2)) - 1
	endIdx := startIdx + b.KernelSize - 1

	startYIdx := y - int(math.Floor(float64(b.KernelSize)/2)) - 1
	endYIdx := startYIdx + b.KernelSize - 1

	kernelInnerSize := b.KernelSize * b.KernelSize

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

	return color.RGBA{
		R: uint8(result.R / kernelInnerSize),
		G: uint8(result.G / kernelInnerSize),
		B: uint8(result.B / kernelInnerSize),
		A: uint8(result.A / kernelInnerSize),
	}
}

func (b *BlurActionMean) Blur() (image.Image, error) {
	var targetImg *image.RGBA
	var pixels [][]Pixel
	var pixelError error

	// iterate over our entire image pixels in blocks of our kernel
	// size. Taking an average of pixel values and applying this
	// back into our new image position.
	for iter := 0; iter < b.Iterations; iter++ {

		if iter == 0 {
			pixels, pixelError = b.validateAndGetImagePixels()
			targetImg = image.NewRGBA(b.Image.Bounds())
		} else {
			b.Image = targetImg
			pixels, pixelError = b.validateAndGetImagePixels()
			targetImg = image.NewRGBA(targetImg.Bounds())
		}

		if pixelError != nil {
			return nil, pixelError
		}

		// IGNORE EDGES
		// TODO: THIS SKIPS EDGES AND WORKS WITH A PERFECT SIZE FIRST, UPDATE TO INCLUDE EDGES
		startingOffset := int(math.Floor(float64(b.KernelSize)/2) + 1)

		var wg sync.WaitGroup

		for i := startingOffset; i < len(pixels)-startingOffset; i++ {
			for j := startingOffset; j < len(pixels[i])-startingOffset; j++ {
				wg.Add(1)

				i := i
				j := j

				go func() {
					// now we must sum all the RGB values within our kernel size
					// and then go and divide this by our total kernel size which
					// would be kernelSize x kernelSize. This would be our new
					// RGB values for the center pixel.
					newPixel := b.getUpdatedPixel(i, j, pixels)
					targetImg.SetRGBA(j, i, newPixel)

					wg.Done()
				}()
			}
		}

		wg.Wait()

	}

	return targetImg, nil
}
