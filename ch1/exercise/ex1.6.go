//@author: hdsfade
//@date: 2020-10-30-09:56
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var red = color.RGBA{0xff,0x00,0x00,0xff}
var green = color.RGBA{0x00,0xff,0x00,0xff}
var blue = color.RGBA{0x00,0x00,0xff,0xff}

var palette = []color.Color{color.White, color.Black, red, green, blue}

const (
	whiteIndex = iota
	blackIndex
	redIndex
	greenIndex
	blueIndex
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func (w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("8080",nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5
		res = 0.001
		nframes = 64
		size = 100
		delay = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i< nframes; i++ {
		rect := image.Rect(0,0,2*size+1,2*size+1)
		img := image.NewPaletted(rect, palette)
		for t:= 0.0; t < cycles*2*math.Pi;t += res {
			x := math.Sin(t)
			y := math.Sin(t *freq + phase)
			colorIndex := rand.Intn(5)
			img.SetColorIndex(size + int(size*x+0.5), size + int(size*y +0.5), uint8(colorIndex))
		}
		phase += 0.1
		anim.Image = append(anim.Image, img)
		anim.Delay = append(anim.Delay, delay)
	}
	gif.EncodeAll(out, &anim)
}