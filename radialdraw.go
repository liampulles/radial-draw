package radialdraw

import (
	"math"

	"github.com/fogleman/gg"
)

// Draw a square image with a radial pattern to the PNG file with name
// given by out. Try it out!
func DrawRadial(out string, width int, iter int) error {
	dc := gg.NewContext(width, width)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	for i := 0.0; i < 2*math.Pi; i += 0.2 {
		drawRay(dc, width, i, 100)
	}
	return dc.SavePNG(out)
}

func drawRay(dc *gg.Context, width int, theta float64, originDisp float64) {
	fWidth := float64(width)
	origin := float64(width) / 2.0
	sinT := math.Sin(theta)
	cosT := math.Cos(theta)
	startX := origin + (originDisp * sinT)
	startY := origin - (originDisp * cosT)
	endX := origin + (fWidth * sinT)
	endY := origin - (fWidth * cosT)

	dc.SetRGBA(1.0, 1.0, 1.0, 1.0)
	dc.SetLineWidth(2.0)
	dc.DrawLine(startX, startY, endX, endY)
	dc.Stroke()
}
