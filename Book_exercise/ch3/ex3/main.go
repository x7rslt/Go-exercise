//给多边形上色，峰顶呈红色，峰低呈蓝色

package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
	defaultColor  = "grey"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

var f1 func(x, y float64) float64

var color string

func main() {
	fmt.Println("http://localhost:8000/?color=gray&f=f")
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			switch k {
			case "color":
				color = v[0]
			case "f":
				switch v[0] {
				case "eggbox":
					f1 = eggbox
				case "saddle":
					f1 = saddle
				default:
					f1 = f
				}
			}
		}
		if color == "" {
			color = defaultColor
		}
		if f1 == nil {
			f1 = f
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		svg(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return

}

func svg(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:%s;fill:white;stroke-width:0.7' "+
		"width='%d' height='%d'>", color, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, fill, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, fill, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, fill, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, fill, ok := corner(i+1, j+1)
			if !ok {
				continue
			}
			fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill:%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, fill)
		}
	}
	fmt.Fprintln(w, "</svg>")
}
func corner(i, j int) (float64, float64, string, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f1(x, y)
	if math.IsInf(z, 0) {
		return 0, 0, "", false
	}
	var fill string
	switch {
	case z > 0.02:
		fill = "#ff0000"
	case z < 0.02:
		fill = "#0000ff"
	default:
		fill = "white"
	}
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, fill, true
}
func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r)
}
func saddle(x, y float64) float64 {
	r := y*y/600 - x*x/300
	return r
}
func eggbox(x, y float64) float64 {
	return (math.Cos(x) + math.Cos(y)) / 10
}
