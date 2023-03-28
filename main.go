package main

import (
	"image/color"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/process-image", Handler)

	e.Static("/", "frontend/dist")

	e.Logger.Fatal(e.Start(":4140"))
}

func getBackground(blackColor bool) color.RGBA {
	if blackColor {
		return color.RGBA{R: 0, G: 0, B: 0, A: 255}
	}

	return color.RGBA{R: 255, G: 255, B: 255, A: 255}
}
