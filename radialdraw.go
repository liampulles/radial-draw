package radialdraw

import (
	"math"

	"github.com/fogleman/gg"
)

// Draw a square image with a radial pattern to the PNG file with name
// given by out. Try it out!
func DrawRadial(out string, width int, spaceFactor float64, lineWidth float64) error {
	// Init and draw background
	dc := gg.NewContext(width, width)
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	// Draw rays
	fullCircle := 2 * math.Pi
	originDisp := 0.0
	fWidth := float64(width)
	maxLen := math.Sqrt(2 * fWidth * fWidth)
	// Keep iterating until we start drawing lines outside the screen
	for i := 0; originDisp < maxLen; i++ {
		rayAngle := fullCircle / math.Pow(2, float64(i+2))
		originDisp = spaceFactor / math.Tan(rayAngle)
		theta := 0.0
		for j := 0; theta < fullCircle; j++ {
			theta += rayAngle
			// Skip previously drawn lines
			if j%2 == 0 {
				continue
			}
			drawRay(dc, width, theta, originDisp, lineWidth)
		}
	}

	// Save
	return dc.SavePNG(out)
}

// Draw a ray from the center out to "infinity", at an angle and displaced by an amount.
func drawRay(dc *gg.Context, width int, theta float64, originDisp float64, lineWidth float64) {
	fWidth := float64(width)
	origin := fWidth / 2.0
	sinT := math.Sin(theta)
	cosT := math.Cos(theta)
	startX := origin + (originDisp * sinT)
	startY := origin - (originDisp * cosT)
	endX := origin + (fWidth * sinT)
	endY := origin - (fWidth * cosT)

	dc.SetRGBA(1.0, 1.0, 1.0, 1.0)
	dc.SetLineWidth(lineWidth)
	dc.DrawLine(startX, startY, endX, endY)
	dc.Stroke()
}
