package imaging

import (
	"github.com/stephensli/image-processing/internal/helpers"
	"image"
	"image/color"
	"math"
	"sync"
)

type BlurActionGaussian struct {
	BlurAction
	kernel [][]float64
	Sigma  float64
}

func (b *BlurActionGaussian) getUpdatedPixel(x, y int, pixels [][]Pixel) color.RGBA {
	startIdx := helpers.Max(x-int(math.Floor(float64(b.KernelSize)/2))-1, 0)
	endIdx := helpers.Min(startIdx+b.KernelSize-1, len(pixels))

	startYIdx := helpers.Max(y-int(math.Floor(float64(b.KernelSize)/2))-1, 0)
	endYIdx := helpers.Min(startYIdx+b.KernelSize-1, len(pixels[0]))

	kernelInnerSize := 0

	result := Pixel{
		R: 0,
		G: 0,
		B: 0,
		A: 0,
	}

	for i := startIdx; i < endIdx; i++ {
		for j := startYIdx; j < endYIdx; j++ {
			// we increment it here since edges will  not have NxN items,
			//so it's easier to have an adjustable value. Otherwise, on the
			// edges it can seem darker.
			kernelInnerSize += 1

			result.R += int(float64(pixels[i][j].R) * b.kernel[i-startIdx][j-startYIdx])
			result.G += int(float64(pixels[i][j].G) * b.kernel[i-startIdx][j-startYIdx])
			result.B += int(float64(pixels[i][j].B) * b.kernel[i-startIdx][j-startYIdx])
			result.A += int(float64(pixels[i][j].A) * b.kernel[i-startIdx][j-startYIdx])
		}
	}

	final := color.RGBA{
		R: uint8(math.Min(float64(result.R/kernelInnerSize), 255)),
		G: uint8(math.Min(float64(result.G/kernelInnerSize), 255)),
		B: uint8(math.Min(float64(result.B/kernelInnerSize), 255)),
		A: uint8(math.Min(float64(result.A/kernelInnerSize), 255)),
	}

	return final
}

func (b *BlurActionGaussian) Blur() (image.Image, error) {
	var targetImg *image.RGBA
	var pixels [][]Pixel
	var pixelError error

	b.kernel = helpers.KernelGaussian(b.KernelSize, b.Sigma)

	// iterate over our entire image pixels in blocks of our kernel
	// size. Use the Gaussian kernel to determine the new value
	// of that given pixel position.

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

		var wg sync.WaitGroup

		for i := 0; i < len(pixels); i++ {
			for j := 0; j < len(pixels[i]); j++ {
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
