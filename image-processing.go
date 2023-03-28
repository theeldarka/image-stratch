package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io"

	"github.com/nfnt/resize"
)

// func openImage(filename string) (image.Image, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	return getImageFromFile(file)
// }

func getImageFromFile(file io.Reader) (image.Image, error) {
	buf, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	img, _, err := image.Decode(bytes.NewReader(buf))
	if err != nil {
		return nil, fmt.Errorf("error decoding image: %w", err)
	}
	return img, nil
}

func resizeImage(img image.Image, width, height int) (image.Image, error) {
	outputRatio := float64(width) / float64(height)
	inputWidth := img.Bounds().Max.X - img.Bounds().Min.X
	inputHeight := img.Bounds().Max.Y - img.Bounds().Min.Y
	inputRatio := float64(inputWidth) / float64(inputHeight)

	var resizeWidth, resizeHeight uint
	if outputRatio > inputRatio {
		resizeWidth = uint(float64(height) * inputRatio)
		resizeHeight = uint(height)
	} else {
		resizeWidth = uint(width)
		resizeHeight = uint(float64(width) / inputRatio)
	}

	resizedImg := resize.Resize(resizeWidth, resizeHeight, img, resize.NearestNeighbor)
	return resizedImg, nil
}

func createImage(width, height int, bgColor color.RGBA) *image.RGBA {
	outputImg := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(outputImg, outputImg.Bounds(), &image.Uniform{C: bgColor}, image.Point{}, draw.Src)
	return outputImg
}

func copyImage(srcImg image.Image, destImg draw.Image) {
	resizeWidth := srcImg.Bounds().Max.X - srcImg.Bounds().Min.X
	resizeHeight := srcImg.Bounds().Max.Y - srcImg.Bounds().Min.Y
	offsetX := (destImg.Bounds().Max.X - resizeWidth) / 2
	offsetY := (destImg.Bounds().Max.Y - resizeHeight) / 2

	// Convert the source image to a draw.Image object
	srcDrawImg := drawImageFromImage(srcImg)

	draw.Draw(destImg, image.Rect(offsetX, offsetY, offsetX+resizeWidth, offsetY+resizeHeight), srcDrawImg, srcDrawImg.Bounds().Min, draw.Src)
}

func drawImageFromImage(img image.Image) draw.Image {
	bounds := img.Bounds()
	dest := image.NewRGBA(bounds)
	draw.Draw(dest, bounds, img, bounds.Min, draw.Src)
	return dest
}

// func saveImage(img image.Image, filename string) error {
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()
// 	err = jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
