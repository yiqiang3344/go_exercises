// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"gopl.io/ch3/surface"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", surface.Width, surface.Height)
	for i := 0; i < surface.Cells; i++ {
		for j := 0; j < surface.Cells; j++ {
			ax, ay, color1, color2 := surface.Corner3(i+1, j)
			bx, by, _, _ := surface.Corner3(i, j)
			cx, cy, _, _ := surface.Corner3(i, j+1)
			dx, dy, _, _ := surface.Corner3(i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='rgb(%g,0,%g)'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color1, color2)
		}
	}
	fmt.Fprintln(w, "</svg>")
}
