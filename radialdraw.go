package radialdraw

import (
	"math"

	"github.com/fogleman/gg"
)

// Draw a square image with a radial pattern to the PNG file with name
// given by out. Try it out!
func DrawRadial(out string, width int, spaceFactor float64) error {
	// Init and draw background
	dc := gg.NewContext(width, width)
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	// Draw rays
	fullCircle := 2 * math.Pi
	originDisp := 0.0
	for i := 0.0; originDisp < float64(width); i++ {
		rayAngle := fullCircle / math.Pow(2, i+2)
		originDisp = spaceFactor / math.Tan(rayAngle)
		// originDisp := 10.0
		for theta := 0.0; theta < fullCircle; theta += rayAngle {
			drawRay(dc, width, theta, originDisp)
		}
	}

	// Save
	return dc.SavePNG(out)
}

// Draw a ray from the center out to "infinity", at an angle and displaced by an amount.
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
