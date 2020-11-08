//@author: hdsfade
//@date: 2020-10-31-10:30
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            //画布大小
	cells         = 100                 //单元格个数
	xyrange       = 30.0                //坐标轴范围
	xyscale       = width / 2 / xyrange //x或y轴单位长度
	zscale        = height * 0.4        //z轴单位长度
	angle         = math.Pi / 6         //x、y轴的角度
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/plot", plot)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func surface(w io.Writer, f func(x, y float64) float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+ //xmlns:命名空间
		"style='stroke:grey;fill:white;stroke-width:0.7' "+ //stroke：线条颜色；fill：填充色；stroke-width：线条宽度
		"width='%d' height='%d'>", width, height) //规定画布大小

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g, %g,%g, %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy) //多边形顶点
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int, f func(x, y float64) float64) (float64, float64) {
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

func parseAndCheck(s string) (Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y)
		return expr.Eval(Env{"x": x, "y": y, "r": r})
	})
}
