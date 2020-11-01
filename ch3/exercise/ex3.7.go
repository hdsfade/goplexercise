//@author: hdsfade
//@date: 2020-10-31-21:01
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	width, height          = 1024, 1024
	eps                    = 1e-6 //精度
	usage                  = "usage: please input gray or color to generate png"
)

type newtons func(complex128) color.Color

func main() {
	var newton newtons
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	} else {
		switch os.Args[1] {
		case "gray":
			newton = newton_gray
		case "color":
			newton = newton_color
		default:
			fmt.Fprintln(os.Stderr, usage)
			os.Exit(1)
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}

//f(z) = z^4 - 1
//f'(z) = 4*z^3
//z' = z- f(z)/f'(z) = z - (z^4 -1) / (4 * z ^3)
//牛顿法求方程近似解

//根据迭代次数灰度上色
func newton_gray(z complex128) color.Color {
	const contrast = 15
	const iteration = 200

	for n := uint8(0); n < iteration; n++ {
		if cmplx.Abs(z*z*z*z-1) < eps {
			return color.Gray{255 - n*contrast}
		}
		z = f(z)
	}

	//迭代次数超过iteration，返回color.Black
	return color.Black
}

//根据求得的根全彩上色
func newton_color(z complex128) color.Color {
	const contrast = 15
	const iteration = 200

	for n := 0; n < iteration; n++ {
		if cmplx.Abs(z*z*z*z-1) < eps {
			change := uint8(cmplx.Abs(z*z*z*z-1) / eps * contrast)
			return color.RGBA{255 + change, change, 255 - change, 255}
		}
		z = f(z)
	}

	//迭代次数超过iteration，返回color.Black
	return color.Black
}

func f(z complex128) complex128 {
	return z - (z*z*z*z-1)/(z*z*z*4)
}
