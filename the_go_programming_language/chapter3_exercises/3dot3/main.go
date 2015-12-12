//Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            //canvas size in pixels
	cells         = 100                 //number of grid cells
	xyrange       = 30.0                //axis ranges (-xyrange ..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        //pixels per z unit
	angle         = math.Pi / 6         //angle of x,y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width ='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)

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

			red := uint8(255 - ((az / 1.2) * 255))
			blue := uint8((az / 1.2) * 255)

			// FF = 15 x 16^1 + 15 * 16^0
			// FF = 240 + 15
			// FF = 255

			//scale 0 to 320 range to 0 -> 255

			//if ay,by,cy,dy are close to 0 it is high if close to 320 low
			// #ff0000 RED peaks
			// #0000ff Blue Valleys

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:#%02X00%02X;' />\n", ax, ay, bx, by, cx, cy, dx, dy, red, blue)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	//Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//Compute surface height z.
	z := f(x, y)

	//Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) //distance from (0,0)
	return math.Sin(r) / r
}
