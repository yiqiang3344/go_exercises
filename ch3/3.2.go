// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"gopl.io/ch3/surface"
)

func main() {
	//for i := 0; i < surface.Cells; i++ {
	//	for j := 0; j < surface.Cells; j++ {
	//		surface.Corner2(i+1, j)
	//	}
	//}
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>", surface.Width, surface.Height)
	for i := 0; i < surface.Cells; i++ {
		for j := 0; j < surface.Cells; j++ {
			ax, ay := surface.Corner2(i+1, j)
			bx, by := surface.Corner2(i, j)
			cx, cy := surface.Corner2(i, j+1)
			dx, dy := surface.Corner2(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}
