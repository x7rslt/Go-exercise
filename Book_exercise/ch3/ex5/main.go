package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
)

func main() {
	createPng()
}

func createPng() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/height*(xmax-xmin) + xmin
			z := complex(x, y)
			// 点(px, py)表示复数值z
			img.Set(px, py, mandelbrot(z))
		}
	}
	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	} // 注意：忽略错误
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			x := 255 - contrast*n
			switch n % 3 {
			case 0:
				return color.RGBA{x, 0, 0, x}
			case 1:
				return color.RGBA{0, x, 0, x}
			case 2:
				return color.RGBA{0, 0, x, x}
			}
		}
	}
	return color.Black
}
