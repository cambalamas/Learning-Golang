package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/Sirupsen/logrus"
)

var (
	fName  = flag.String("n", curPath(), "Name of the image file.")
	fSize  = flag.Int("s", 250, "Define the image file size.")
	fColor = flag.Int("c", 255, "Define the limit of bright 3 - 255.")
	fAlpha = flag.Int("a", 255, "Define the alpha color 0 - 255.")
)

func main() {
	flag.Parse()
	chkColor(fAlpha)
	chkColor(fColor)
	rand.Seed(time.Now().Unix())
	canvas := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{*fSize, *fSize}})

	// Fill the canvas with random data == noise
	for x := 0; x < *fSize; x++ {
		for y := 0; y < *fSize; y++ {
			c := color.RGBA{randColor(), randColor(), randColor(), uint8(*fAlpha)}
			canvas.Set(x, y, c)
		}
	}
	// Image file where we deserve to save that random data
	imgFile, err := os.Create(*fName + ".png")
	if err != nil {
		logrus.Warn(err)
	}
	// Dump random data into file
	if err := png.Encode(imgFile, canvas); err != nil {
		logrus.Warn(err)
	}
}

// Helpers ...

func curPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		logrus.Warn(err)
	}
	return filepath.Base(cwd)
}

func chkColor(cv *int) {
	switch {
	case *cv > 255:
		*cv = 255
	case *cv < 3:
		*cv = 3
	}
}

func randColor() uint8 {
	return uint8(rand.Intn(*fColor))
}
