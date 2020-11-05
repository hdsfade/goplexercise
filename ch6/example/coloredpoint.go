//@author: hdsfade
//@date: 2020-11-05-20:26
package point

import (
	"image/color"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}
