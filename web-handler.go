package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sunshineplan/imgconv"
)

// type Color struct {
// 	Red   int `form:"color.red" validate:"required,gte=0,lte=255"`
// 	Green int `form:"color.green" validate:"required,gte=0,lte=255"`
// 	Blue  int `form:"color.blue" validate:"required,gte=0,lte=255"`
// }

type ProcessImageRequest struct {
	Image  *multipart.FileHeader `form:"image"`
	Width  int                   `form:"width" validate:"required"`
	Height int                   `form:"height" validate:"required"`
	// Color  Color                 `form:"color" validate:"required"`
	BlackBackground bool `form:"black_background"`
}

type CustomValidator struct {
	validator *validator.Validate
}

type ErrorResponse struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Handler(c echo.Context) error {
	req := new(ProcessImageRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request parameters",
		})
	}

	if err := c.Validate(req); err != nil {
		errors := make(map[string]string, 0)
		for _, e := range err.(validator.ValidationErrors) {

			errors[strings.Replace(e.Namespace(), "ProcessImageRequest.", "", 1)] = e.Tag()
		}
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "Validation error",
			Errors:  errors,
		})
	}

	inputImg, err := getImageFromRequest(c)
	if err != nil {
		log.Error(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to parse image: %s", err.Error()),
		})
	}

	bgColor := getBackground(req.BlackBackground)
	outputWidth := req.Width
	outputHeight := req.Height

	resizedImg, err := resizeImage(inputImg, outputWidth, outputHeight)
	if err != nil {
		log.Error(err)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to resize image: %s", err.Error()),
		})
	}

	outputImg := createImage(outputWidth, outputHeight, bgColor)

	copyImage(resizedImg, outputImg)

	c.Response().Header().Set("Content-Disposition", "attachment; filename=image.jpg")
	c.Response().Header().Set("Content-Type", "image/jpeg")

	if err := jpeg.Encode(c.Response().Writer, outputImg, &jpeg.Options{Quality: 100}); err != nil {
		log.Error(err.Error)

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Failed to encode result image: %s", err.Error()),
		})
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

	return imgconv.Decode(file)
}
