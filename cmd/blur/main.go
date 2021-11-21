package main

import (
	"log"
	"os"

	localCli "github.com/stephensli/image-processing/internal/cli"
	"github.com/urfave/cli/v2"
)

func main() {
	var app = &cli.App{
		Name:  "Image Processing",
		Usage: "Perform Basic Image Processing",
		Commands: []*cli.Command{
			{
				Name:   "blur",
				Usage:  "Perform a blur on a given image",
				Action: localCli.PerformBlurOnImage,
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:     "file",
						Aliases:  []string{"f"},
						Usage:    "The path to the file being blurred.",
						Required: true,
					},
					&cli.StringSliceFlag{
						Name:        "type",
						Aliases:     []string{"t"},
						Usage:       "The type of blur to apply",
						Required:    true,
						Value:       cli.NewStringSlice("mean", "gaussian"),
						DefaultText: "mean",
					},
					&cli.IntFlag{
						Name:    "kernel",
						Aliases: []string{"k"},
						Usage:   "The size of the kernel used on the blur",
						Value:   31,
					},
					&cli.Float64Flag{
						Name:    "sigma",
						Aliases: []string{"s"},
						Usage:   "Sets the sigma value if used in the blur, e.g Gaussian blur.",
						Value:   20,
					},
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
