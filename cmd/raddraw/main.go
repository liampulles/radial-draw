package main

import (
	"flag"
	"fmt"
	"os"

	radialdraw "github.com/liampulles/radial-draw"
)

func main() {
	args := parseArgs()
	if err := radialdraw.DrawRadial(args.Out, args.Width, args.Space, args.Line); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
		os.Exit(1)
	}
}

type args struct {
	Out   string
	Width int
	Space float64
	Line  float64
}

func parseArgs() args {
	out := flag.String("out", "radial.png", "Location of PNG to write")
	width := flag.Int("width", 1024, "width of the image (square dimensions)")
	space := flag.Float64("space", 5.0, "spacing factor")
	line := flag.Float64("line", 2.0, "line width")
	flag.Parse()
	return args{
		Out:   *out,
		Width: *width,
		Space: *space,
		Line:  *line,
	}
}
