package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	width := 320
	height := 320
	green := color.RGBA{0, 255, 0, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)
	for i := 0; i < 320; i += 40 {
		for y := 0; y < 320; y++ {
			rectImg.Set(i, y, color.Black)
		}
	}
	for i := 0; i < 320; i += 40 {
		for y := 0; y < 320; y++ {
			rectImg.Set(y, i, color.Black)
		}
	}

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, rectImg)
}
