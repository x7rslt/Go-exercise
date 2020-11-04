//修改利萨茹服务器以通过URL参数读取参数值。例如，你可以通过调整它，使得像http://localhost:8000/?cycle=20这样的网址将其周期设置为20，
//以替代默认的5.使用strconv.Atoi函数将字符串参数转为整型。可以通过go doc strconv.Atoi来查看文档
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

import (
	"log"
	"net/http"
	"time"
	"net/url"
	"strconv"
	"fmt"
)

//!+main

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	//cycles := 5
	rand.Seed(time.Now().UTC().UnixNano())
	handler := func(w http.ResponseWriter, r *http.Request) {
		q,_ := url.ParseQuery(r.URL.RawQuery)
		if _,ok := q["cycles"];ok{
			value,_ := strconv.ParseFloat(q["cycles"][0],64)
			lissajous(w,value)
		}else{
			fmt.Fprintf(w,"%s\n","Error please input cycles value!")
		}
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer,value float64) {
	var cycles float64
	cycles = value
	const (
		//cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
