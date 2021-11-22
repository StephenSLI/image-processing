package helpers

import (
	"math"
)

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// Max returns the greater of x or y.
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Gaussian(x float64, s float64) float64 {
	sig := 2 * s * s
	return math.Exp(-(math.Pow(x, 2) / sig))
}

func KernelGaussian(kernelSize int, sig float64) [][]float64 {
	arr := ARangeAutoStep(0, kernelSize)

	for index, val := range arr {
		arr[index] = Gaussian(val, sig)
	}

	kernel := make([][]float64, kernelSize)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			kernel[i] = append(kernel[i], arr[i]*arr[j])
		}
	}

	return kernel
}

func ARangeAutoStep(start, stop int) []float64 {
	rnge := make([]float64, stop, stop+1)
	step := (float64(stop) - 1.0) / 2.0

	for i := start; i < stop; i++ {
		rnge[i] = float64(i) - step

	}
	return rnge
}
