//@author: hdsfade
//@date: 2020-10-31-16:47
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2 //复数范围
		width, height          = 1024, 1024   //画布大小
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height)) //创建彩色画布

	//1024*1024灰度图迭代
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbort(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbort(z complex128) color.Color {
	const iteration = 200 //迭代次数，超过此次数即视为迭代了无穷次
	const contrast = 15   //对比度

	var v complex128

	for n := uint8(0); n < iteration; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - n*contrast}
		}
	}
	return color.Black
}
