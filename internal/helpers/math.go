package helpers

import (
	"fmt"
	"math"
)

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Gaussian(x float64, s float64) float64 {
	//sig := 2 * s * s

	//return math.Exp(-(math.Pow(x, 2) / sig))

	return math.Pi * (1.0 / s) * math.Exp(-0.5*(math.Pow(x, 2.0)/math.Pow(s, 2.0)))
}

func KernelGaussian(kernelSize int) [][]float64 {
	arr := ARangeAutoStep(0, kernelSize)

	fmt.Println("before", arr)
	for index, val := range arr {
		arr[index] = Gaussian(val, 3)
	}

	kernel := make([][]float64, kernelSize)
	fmt.Println("after", arr)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			kernel[i] = append(kernel[i], arr[i]*arr[j])
		}
	}

	fmt.Println("guassian matrix", kernel)

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
