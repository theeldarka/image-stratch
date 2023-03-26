package main

import (
	"image"
	"image/jpeg"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Handler(c echo.Context) error {
	inputImg, err := getImageFromRequest(c)
	if err != nil {
		log.Info(err)

		return err
	}

	outputWidth, err := strconv.Atoi(c.FormValue("width"))
	if err != nil {
		log.Infof("Failed to parse width %s", c.FormValue("width"))

		outputWidth = 1920
	}

	outputHeight, err := strconv.Atoi(c.FormValue("height"))
	if err != nil {
		log.Infof("Failed to parse height %s", c.FormValue("height"))

		outputHeight = 1080
	}

	wantBlackBackground, err := strconv.ParseBool(c.FormValue("black_background"))
	if err != nil {
		log.Infof("Failed to black_background %s", c.FormValue("black_background"))

		wantBlackBackground = true
	}

	bgColor := getBackground(wantBlackBackground)

	resizedImg, err := resizeImage(inputImg, outputWidth, outputHeight)
	if err != nil {
		log.Error(err.Error)

		return err
	}

	outputImg := createImage(outputWidth, outputHeight, bgColor)

	copyImage(resizedImg, outputImg)

	c.Response().Header().Set("Content-Disposition", "attachment; filename=image.jpg")
	c.Response().Header().Set("Content-Type", "image/jpeg")

	if err := jpeg.Encode(c.Response().Writer, outputImg, &jpeg.Options{Quality: 100}); err != nil {
		log.Error(err.Error)

		return err
	}

	return nil
}

func getImageFromRequest(c echo.Context) (image.Image, error) {
	form, err := c.FormFile("image")
	if err != nil {
		return nil, err
	}

	file, err := form.Open()
	if err != nil {
		return nil, err
	}

	defer file.Close()

	inputImg, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}

	return inputImg, nil
}
