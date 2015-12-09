package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{0x5B, 0xB9, 0x53, 0xFF}, color.RGBA{0x4A, 0xB3, 0xF0, 0xFF}, color.RGBA{0xEC, 0x59, 0x35, 0xFF}, color.RGBA{0xDE, 0x48, 0x42, 0xFF}}

const (
	whiteIndex  = 0
	blackIndex  = 1
	greenIndex  = 2
	blueIndex   = 3
	orangeIndex = 4
	redIndex    = 5
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	color := uint8(r1.Intn(redIndex + 1))
	lissajous(os.Stdout, color)
}

func lissajous(out io.Writer, color uint8) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), color)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
