//@author: hdsfade
//@date: 2020-10-31-20:36
package main

import (
	"image"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	width, height          = 1024, 1024
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			ax, ay := corner(px, py+1)
			bx, by := corner(px, py+1)
			cx, cy := corner(px+1, py+1)
			dx, dy := corner(px+1, py+1)

			az, bz := complex(ax, ay), complex(bx, by)
			cz, dz := complex(cx, cy), complex(dx, dy)

			colour := (mandelbrot(az) + mandelbrot(bz) + mandelbrot(cz) + mandelbrot(dz)) / 4
			img.Set(px, py, color.Gray{colour})
		}
	}
	png.Encode(os.Stdout, img)
}

func corner(px, py int) (x, y float64) {
	x = float64(px)/width*(xmax-xmin) + xmin
	y = float64(py)/height*(ymax-ymin) + ymin
	return x, y
}

func mandelbrot(z complex128) uint8 {
	const iteration = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iteration; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return 255 - contrast*n
		}
	}
	return 0
}
