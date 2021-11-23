package imaging

import (
	"image"
	"image/color"
	"math"
	"sync"

	"github.com/stephensli/image-processing/internal/helpers"
)

type BlurActionGaussian struct {
	BlurAction
	kernel [][]float64
	Sigma  float64
}

// kernelGaussian will generate a new kernel to be used in determining the new
// pixel value. Each kernel value be based on the gaussian function with a shift
// based on the sigma value.
//
// https://en.wikipedia.org/wiki/Gaussian_function
func kernelGaussian(kernelSize int, sig float64) [][]float64 {
	arr := helpers.ARangeAutoStep(0, kernelSize)

	for index, val := range arr {
		arr[index] = helpers.Gaussian(val, sig)
	}

	kernel := make([][]float64, kernelSize)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			kernel[i] = append(kernel[i], arr[i]*arr[j])
		}
	}

	return kernel
}

// getUpdatedPixel returns a new color.RGBA for the given pixel location provided
// at the x, y position. The updated pixel will be based on the gaussian kernel.
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

	pixels, pixelError := b.validateAndGetImagePixels()

	// determine the kernel gaussian which will be used to determine the new
	// pixel value by applying a gaussian weighting, the center pixel will
	// contain the highest weighting, decreasing outwards.
	b.kernel = kernelGaussian(b.KernelSize, b.Sigma)

	for iter := 0; iter < b.Iterations; iter++ {
		targetImg = image.NewRGBA(b.Image.Bounds())

		if pixelError != nil {
			return nil, pixelError
		}

		var wg sync.WaitGroup

		// iterate over each pixel within the image and determine the new pixel value.
		// Once the new pixel value is determined, update the new target image pixel
		// location.
		for i := 0; i < len(pixels); i++ {
			for j := 0; j < len(pixels[i]); j++ {
				wg.Add(1)

				go func(i, j int, pixels [][]Pixel) {
					newPixel := b.getUpdatedPixel(i, j, pixels)
					targetImg.SetRGBA(j, i, newPixel)

					wg.Done()
				}(i, j, pixels)
			}
		}

		wg.Wait()

		b.Image = targetImg
		pixels, pixelError = b.validateAndGetImagePixels()

	}

	return targetImg, nil
}
