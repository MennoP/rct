// Draw a RCT roller coaster in 2d.
package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"

	"code.google.com/p/draw2d/draw2d"
	rct "github.com/kevinburke/rct-rides"
)

func saveToPngFile(filePath string, m image.Image) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, m)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s OK.\n", filePath)
}

const PIECE_WIDTH = 8
const IMG_HEIGHT = 400

var RED = color.RGBA{0xff, 0x00, 0x00, 0xff}
var BLUE = color.RGBA{0x00, 0x00, 0xff, 0xff}

func Draw(r *rct.Ride) image.Image {
	width := PIECE_WIDTH * (len(r.TrackData.Elements) + 2)
	rect := image.Rect(0, 0, width, IMG_HEIGHT)
	i := image.NewRGBA(rect)
	c := color.RGBA{0xff, 0xff, 0xff, 0xff}
	draw.Draw(i, i.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	gc := draw2d.NewGraphicContext(i)
	x := float64(PIECE_WIDTH)
	y := float64(IMG_HEIGHT) - 20.0
	for j := 0; j < len(r.TrackData.Elements); j++ {
		gc.MoveTo(x, y)
		elem := r.TrackData.Elements[j]
		seg := elem.Segment
		if elem.ChainLift {
			gc.SetStrokeColor(BLUE)
		} else {
			gc.SetStrokeColor(RED)
		}
		y -= 10.0 * float64(seg.ElevationDelta)
		gc.LineTo(float64(x+PIECE_WIDTH), y)
		gc.Stroke()
		x += PIECE_WIDTH
	}
	return i
}

// 0 to -1 ==> -1 (diff -1)
// -1 to 0 ==> 0 (diff 1)

// 0 to 1 ==> 0 (diff 1)
// 1 to 0 ==> 1 (diff -1)

// 1 to 1 ==> 1
// 4 to 4 ==> 4
// 1 to 4 ==> 2
// 4 to 1 ==> 2
// 0-1, 1-4, 4-4, 4-1, 1-0

func main() {
	r := rct.ReadRide("../rides/mischief.td6")
	img := Draw(r)
	saveToPngFile("test.png", img)
}
