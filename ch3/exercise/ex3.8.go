//@author: hdsfade
//@date: 2020-10-31-22:48
//float64和bigrat放大比较
package main

import (
	"image"
	"image/png"
	"math/cmplx"
	"os"
	"image/color"
)

const (
	magnify = 20 //放大倍数
	xmin, ymin, xmax, ymax = -2*magnify, -2*magnify, 2*magnify, 2*magnify
	width, height          = 1024*magnify, 1024*magnify
	eps                    = 1e-6 //精度
	usage                  = "usage: please input gray or color to generate png"
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, gray(z))
		}
	}
	png.Encode(os.Stdout, img)
}

//f(z) = z^4 - 1
//f'(z) = 4*z^3
//z' = z- f(z)/f'(z) = z - (z^4 -1) / (4 * z ^3)
//牛顿法求方程近似解

//根据迭代次数灰度上色
func gray(z complex128) color.Color {
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

func f(z complex128) complex128 {
	return z - (z*z*z*z-1)/(z*z*z*4)
}
