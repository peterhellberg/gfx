package main

import "github.com/peterhellberg/gfx"

func main() {
	m := gfx.NewImage(64, 64, gfx.ColorWhite)

	gfx.DrawLineBresenham(m, gfx.V(10, 10), gfx.V(54, 54), gfx.ColorRed)
	gfx.DrawLineBresenham(m, gfx.V(10, 20), gfx.V(10, 54), gfx.ColorGreen)
	gfx.DrawLineBresenham(m, gfx.V(20, 10), gfx.V(54, 10), gfx.ColorBlue)

	s := gfx.NewScaledImage(m, 4)

	gfx.SavePNG("/tmp/gfx-readme-examples-bresenham-line.png", s)
}
