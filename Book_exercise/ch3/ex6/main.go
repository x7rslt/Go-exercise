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
	var colors [height][height]color.Color
	img := image.NewRGBA(image.Rect(0, 0, width/2, height/2))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/height*(xmax-xmin) + xmin
			z := complex(x, y)
			// 点(px, py)表示复数值z
			//img.Set(px, py, mandelbrot(z))
			colors[py][px] = mandelbrot(z)
		}
	}
	for py2 := 0; py2 < height/2; py2++ {
		for px2 := 0; px2 < width/2; px2++ {
			c1 := colors[2*py2][2*px2]
			c1r, c1g, c1b, c1a := c1.RGBA()
			c2 := colors[2*py2+1][2*px2]
			c2r, c2g, c2b, c2a := c2.RGBA()
			c3 := colors[2*py2][2*px2+1]
			c3r, c3g, c3b, c3a := c3.RGBA()
			c4 := colors[2*py2+1][2*px2+1]
			c4r, c4g, c4b, c4a := c4.RGBA()
			cr := uint8((c1r + c2r + c3r + c4r) / 4)
			cg := uint8((c1g + c2g + c3g + c4g) / 4)
			cb := uint8((c1b + c2b + c3b + c4b) / 4)
			ca := uint8((c1a + c2a + c3a + c4a) / 4)
			img.Set(px2, py2, color.RGBA{cr, cg, cb, ca})
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
