//@author: hdsfade
//@date: 2020-10-31-19:29
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	width, height          = 1024, 1024
)

func main() {
	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const (
		iteration = 200
		contrast  = 15
	)
	var v complex128
	for n := uint8(0); n < iteration; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			change := n * contrast
			return color.RGBA{255 + change, change, 255 - change, 255}
		}
	}
	return color.Black
}
