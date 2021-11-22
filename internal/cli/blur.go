package cli

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/stephensli/image-processing/internal/imaging"
	"github.com/urfave/cli/v2"
)

func PerformBlurOnImage(c *cli.Context) error {
	fullFilePath := c.String("file")
	blurKind := c.StringSlice("type")[0]

	// validate that the file path exists and that the given file  is one of
	// the possible valid file extensions.
	if _, err := os.Stat(fullFilePath); errors.Is(err, os.ErrNotExist) {
		return errors.New(fmt.Sprintf("file does not exist by path %s", fullFilePath))
	}

	fileExt := strings.ToLower(filepath.Ext(fullFilePath))
	if pos := strings.Index(strings.Join([]string{".jpg", ".jpeg", ".png"}, ""), fileExt); pos == -1 {
		return errors.New(fmt.Sprintf("file extension not supported, extension %s", fileExt))
	}

	file, _ := os.Open(fullFilePath)
	defer file.Close()

	var blurAction imaging.Blur

	kernelSize := c.Int("kernel")
	iterations := c.Int("iterations")

	img, _, _ := image.Decode(file)

	blurAct := imaging.BlurAction{
		KernelSize: kernelSize,
		Iterations: iterations,
		Image:      img,
	}

	if blurKind == "mean" {
		blurAction = &imaging.BlurActionMean{BlurAction: blurAct}
	}

	if blurKind == "gaussian" {
		blurAction = &imaging.BlurActionGaussian{
			Sigma:      c.Float64("sigma"),
			BlurAction: blurAct,
		}
	}

	targetImg, blurringError := blurAction.Blur()

	if blurringError != nil {
		return blurringError
	}

	folderPath := filepath.Dir(fullFilePath)
	fileName := strings.Split(filepath.Base(fullFilePath), ".")[0]

	outputImg, _ := os.Create(fmt.Sprintf("%s/%s-blur%s", folderPath, fileName, fileExt))
	defer outputImg.Close()

	var encodingError error

	switch fileExt {
	case ".jpeg":
	case ".jpg":
		encodingError = jpeg.Encode(outputImg, targetImg, nil)
		break
	case ".png":
		encodingError = png.Encode(outputImg, targetImg)
		break
	}

	return encodingError
}
