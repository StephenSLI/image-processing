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

	targetImg, err := imaging.BlurMean(file, 26)

	f, _ := os.Create("C:\\Users\\stephen\\Pictures\\dog-2.png")
	defer f.Close()

	png.Encode(f, targetImg)

	if err != nil {
		log.Fatalln(err)
	}
}
