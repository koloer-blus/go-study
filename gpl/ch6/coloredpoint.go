package main

import "image/color"

type Points struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	cp.Y = 2
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
}
