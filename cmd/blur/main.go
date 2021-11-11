package main

import (
	"github.com/mambadev/image-processing/internal/imaging"
	"image"
	"image/png"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()

	sourceName := "dog"

	// You can register another format here
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	file, _ := os.Open("C:\\Users\\stephen\\Pictures\\" + sourceName + ".png")
	fileTwo, _ := os.Open("C:\\Users\\stephen\\Pictures\\" + sourceName + ".png")

	defer file.Close()
	defer fileTwo.Close()

	targetImg, _ := imaging.BlurGaussian(file, 31, 20)
	targetImgMean, _ := imaging.BlurMean(fileTwo, 21)

	first, _ := os.Create("C:\\Users\\stephen\\Pictures\\" + sourceName + "-g-blur.png")
	defer first.Close()

	png.Encode(first, targetImg)

	second, _ := os.Create("C:\\Users\\stephen\\Pictures\\" + sourceName + "-mean-blur.png")
	defer second.Close()

	png.Encode(second, targetImgMean)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
