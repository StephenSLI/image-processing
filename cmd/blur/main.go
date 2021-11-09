package main

import (
	"github.com/mambadev/image-processing/internal/imaging"
	"image"
	"image/png"
	"log"
	"os"
)

func main() {
	// You can register another format here
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	file, _ := os.Open("C:\\Users\\stephen\\Pictures\\dog.png")

	defer file.Close()

	err := imaging.BlurGaussian(file, 7)

	if err != nil {
		log.Fatalln(err)
	}
}
