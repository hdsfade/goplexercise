//@author: hdsfade
//@date: 2020-10-31-12:34
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	//"strconv"
)

const (
	cells   = 100         //单元格个数
	xyrange = 30.0        //坐标轴范围
	angle   = math.Pi / 6 //x、y轴的角度
)

var (
	width, height = 600, 320 //画布大小
	color         = "white"  //颜色
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", svg)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func svg(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.FormValue("width") != "" {
		width, err = strconv.Atoi(r.FormValue("width"))
	}
	if err != nil {
		width = 600
	}

	if r.FormValue("height") != "" {
		height, err = strconv.Atoi(r.FormValue("height"))
	}
	if err != nil {
		height = 320
	}

	color = r.FormValue("color")
	if color == "" {
		color = "white"
	}

	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+ //xmlns:命名空间
		"style='stroke:grey;fill:%s;stroke-width:0.7' "+ //stroke：线条颜色；fill：填充色；stroke-width：线条宽度
		"width='%d' height='%d'>", color, width, height) //规定画布大小

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			//函数f返回无穷大值，只可能引起y坐标无穷大
			if math.IsInf(ay, 0) || math.IsInf(by, 0) || math.IsInf(cy, 0) || math.IsInf(dy, 0) {
				continue
			}

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g, %g,%g, %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy) //多边形顶点
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int) (float64, float64) {
	xyscale := float64(width) * 0.5 / xyrange //x或y轴单位长度
	zscale := float64(height) * 0.4           //z轴单位长度

	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//计算曲面高度z
	//可根据不同曲面高度函数，绘制不同图形
	z := f(x, y)

	//将(x,y,z)等角投影到二维平面上
	//对于公式原理有兴趣的，可深入了解
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
