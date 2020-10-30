//@author: hdsfade
//@date: 2020-10-30-14:36
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)
var red = color.RGBA{0xff,0x00,0x00,0xff}
var green = color.RGBA{0x00,0xff,0x00,0xff}
var blue = color.RGBA{0x00,0x00,0xff,0xff}

var paletee = []color.Color{color.White, color.Black, red, green, blue}

/*const (
	whiteIndex = iota
	blackIndex
	redIndex
	greenIndex
	blueIndex
)*/

var colors = map[string]int{"white":0, "black":1,"red":2,"green":3,"blue":4}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	lissajous(w, r)
}

func lissajous(out http.ResponseWriter, r *http.Request) {
	const size = 100
	var err error
	var (
		cycles  = 5.0
		res     = 0.001
		nframes = 64
		delay   = 8
		color = "black"
	)

	r.ParseForm()
	for k, v := range r.Form {
		switch k {
		case "cycles":
			cycles, err= strconv.ParseFloat(v[0], 64)
			checkErr(out, err)
		case "res":
			res, _ = strconv.ParseFloat(v[0], 64)
			checkErr(out, err)
		case "nframes":
			nframes, _ = strconv.Atoi(v[0])
			checkErr(out, err)
		case "delay":
			delay, _ = strconv.Atoi(v[0])
			checkErr(out, err)
		case "color":
			color = v[0]
		}
	}
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, paletee)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(size*x+0.5), size+int(size*y+0.5), uint8(colors[color]))
		}
		phase += 0.1
		anim.Image = append(anim.Image, img)
		anim.Delay = append(anim.Delay, delay)
	}
	gif.EncodeAll(out, &anim)
}

func checkErr(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Fprintf(w, "set paraments: %v\nuse default\n", err)
	}
}
