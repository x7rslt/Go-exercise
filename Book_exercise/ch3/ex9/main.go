package main

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
)

func main() {
	fmt.Println("http://localhost:8000")
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `<body>`)
		fmt.Fprint(w, `<img src="data:image/png;base64,`)
		createPng(w)
		fmt.Fprint(w, `" />`)
		fmt.Fprint(w, `</body>`)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func createPng(w http.ResponseWriter) {
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
	b64w := base64.NewEncoder(base64.StdEncoding, w) // 往b64w里写，就是编码后写入到w
	defer b64w.Close()
	png.Encode(b64w, img) // 注意：忽略错误
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
