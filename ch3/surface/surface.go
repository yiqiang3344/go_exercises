// Surface computes an SVG rendering of a 3-D surface function.
package surface

import (
	"math"
)

const (
	Width, Height = 600, 320            // canvas size in pixels
	Cells         = 100                 // number of grid cells
	Xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	Xyscale       = Width / 2 / Xyrange // pixels per x or y unit
	Zscale        = Height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var Sin30, Cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func Corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := Xyrange * (float64(i)/Cells - 0.5)
	y := Xyrange * (float64(j)/Cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := Width/2 + (x-y)*Cos30*Xyscale
	sy := Height/2 + (x+y)*Sin30*Xyscale - z*Zscale
	return sx, sy
}

func Corner1(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := Xyrange * (float64(i)/Cells - 0.5)
	y := Xyrange * (float64(j)/Cells - 0.5)

	// Compute surface height z.
	z := f1(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := Width/2 + (x-y)*Cos30*Xyscale
	sy := Height/2 + (x+y)*Sin30*Xyscale - z*Zscale
	return sx, sy
}

func Corner2(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := Xyrange * (float64(i)/Cells - 0.5)
	y := Xyrange * (float64(j)/Cells - 0.5)

	// Compute surface height z.
	z := f2(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := Width/2 + (x-y)*Cos30*Xyscale
	sy := Height/2 + (x+y)*Sin30*Xyscale - z*Zscale
	return sx, sy
}

func Corner3(i, j int) (float64, float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := Xyrange * (float64(i)/Cells - 0.5)
	y := Xyrange * (float64(j)/Cells - 0.5)

	// Compute surface height z.
	z, color1, color2 := f3(x, y)

	//fmt.Printf("\n %g %g %g\n", x, y, z)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := Width/2 + (x-y)*Cos30*Xyscale
	sy := Height/2 + (x+y)*Sin30*Xyscale - z*Zscale

	return sx, sy, color1, color2
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func f1(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	if r == 0 {
		return 0
	}
	return math.Sin(r) / r
}

func f2(x, y float64) float64 {
	if x == 0 && y == 0 {
		return 0
	}
	var v1, v2, v3, v4 float64 = 1, 4, 1, 10
	return (v3/v4*math.Pow(y, 2) - v1/v2*math.Pow(x, 2))/13
}

func f3(x, y float64) (float64, float64, float64) {
	r := math.Hypot(x, y) // distance from (0,0)
	if r == 0 {
		return 0, 0, 0
	}
	color1 := (math.Sin(r) + 1) / 2 * 255
	color2 := (1 - math.Sin(r) + 1) / 2 * 255
	return math.Sin(r) / r, color1, color2
}
