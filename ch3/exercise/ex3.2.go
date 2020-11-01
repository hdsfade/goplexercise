//@author: hdsfade
//@date: 2020-10-31-12:07
package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

//定义生成曲面高度z的函数类型
type zFunc func(x, y float64) float64

const (
	width, height = 600, 320            //画布大小
	cells         = 100                 //单元格个数
	xyrange       = 30.0                //坐标轴范围
	xyscale       = width / 2 / xyrange //x或y轴单位长度
	zscale        = height * 0.4        //z轴单位长度
	angle         = math.Pi / 6         //x、y轴的角度
)

//用法提示
const usage = "usage: please input saddle or eggbox to generate svg"

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	var f zFunc

	switch os.Args[1] {
	case "eggbox":
		f = eggbox
	case "saddle":
		f = saddle
	default:
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	svg(os.Stdout, f)
}

//根据zFunc产生svg图像并输出到out中
func svg(out io.Writer, f zFunc) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+ //xmlns:命名空间
		"style='stroke:grey;fill:white;stroke-width:0.7' "+ //stroke：线条颜色；fill：填充色；stroke-width：线条宽度
		"width='%d' height='%d'>", width, height) //规定画布大小

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j,f)
			bx, by := corner(i, j,f)
			cx, cy := corner(i, j+1,f)
			dx, dy := corner(i+1, j+1,f)
			fmt.Fprintf(out,"<polygon points='%g,%g %g,%g, %g,%g, %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy) //多边形顶点
		}
	}
	fmt.Fprintln(out, "</svg>")
}

//生成二维坐标图
func corner(i, j int, f zFunc) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//计算曲面高度z
	//可根据不同曲面高度函数，绘制不同图形
	z := f(x, y)

	//将(x,y,z)等角投影到二维平面上
	//对于公式原理有兴趣的，可深入了解
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

//鸡蛋盒生成函数
func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

//马鞍生成函数
func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	return y*y/a2 - x*x/b2
}