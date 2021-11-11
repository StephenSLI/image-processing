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

	// You can register another format here
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	file, _ := os.Open("C:\\Users\\stephen\\Pictures\\dog.png")

	defer file.Close()

	targetImg, err := imaging.BlurGaussian(file, 31)

	if err != nil {
		log.Fatalln(err)
	}

	f, _ := os.Create("C:\\Users\\stephen\\Pictures\\dog-blur.png")
	defer f.Close()

	png.Encode(f, targetImg)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
