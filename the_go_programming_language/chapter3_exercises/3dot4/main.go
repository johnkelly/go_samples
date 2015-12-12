//Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         //number of grid cells
	xyrange = 30.0        //axis ranges (-xyrange ..+xyrange)
	angle   = math.Pi / 6 //angle of x,y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/graph", graphHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func graphHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	height := r.URL.Query().Get("height")
	width := r.URL.Query().Get("width")

	if height == "" {
		height = "320"
	}

	if width == "" {
		width = "600"
	}

	numHeight, _ := strconv.Atoi(height)
	numWidth, _ := strconv.Atoi(width)

	constructGraph(w, numHeight, numWidth)
}

func constructGraph(w io.Writer, height int, width int) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width ='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j, height, width)
			bx, by, _ := corner(i, j, height, width)
			cx, cy, _ := corner(i, j+1, height, width)
			dx, dy, _ := corner(i+1, j+1, height, width)

			values := []float64{ax, ay, bx, by, cx, cy, dx, dy}
			printable := true

			for _, value := range values {
				if math.IsNaN(value) || math.IsInf(value, 0) {
					printable = false
					break
				}
			}

			if printable == false {
				continue
			}

			red := uint8(255 - ((az / 5) * 255))
			blue := uint8((az / 5) * 255)

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:#%02X00%02X;' />\n", ax, ay, bx, by, cx, cy, dx, dy, red, blue)
		}
	}
	fmt.Fprint(w, "</svg>")
}

func corner(i, j, height, width int) (float64, float64, float64) {
	//Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//Compute surface height z.
	z := f(x, y)
	zscale := float64(height) * 0.4         //pixels per z unit
	xyscale := float64(width) / 2 / xyrange // pixels per x or y unit

	//Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) //distance from (0,0)
	return math.Sin(r) / r
}
