package main

import (
	"flag"
	"fmt"
	"os"

	radialdraw "github.com/liampulles/radial-draw"
)

func main() {
	args := parseArgs()
	if err := radialdraw.DrawRadial(args.Out, args.Width, args.Iter); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
		os.Exit(1)
	}
}

type args struct {
	Out   string
	Width int
	Iter  int
}

func parseArgs() args {
	out := flag.String("out", "radial.png", "Location of PNG to write")
	width := flag.Int("width", 1024, "width of the image (square dimensions)")
	iter := flag.Int("iter", 10, "number of iterations")
	return args{
		Out:   *out,
		Width: *width,
		Iter:  *iter,
	}
}
