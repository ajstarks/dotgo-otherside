package main

import (
	"github.com/ajstarks/svgo"
	"os"
)

func main() {
	width := 960
	height := 540
	style := "fill:white;font-size:60pt;font-family:serif;text-anchor:middle"
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:black")
	canvas.Circle(width/2, height, width/2, "fill:rgb(44,77,232)")
	canvas.Text(width/2, height/2, "hello, world", style)
	canvas.End()
}
