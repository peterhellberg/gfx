package main

import "github.com/peterhellberg/gfx"

func main() {
	m := gfx.NewImage(512, 512, gfx.ColorWhite)

	gfx.NewTurtle(gfx.V(148, 450), func(t *gfx.Turtle) {
		t.Color = gfx.ColorWithAlpha(gfx.ColorMagenta, 64)

		for i := 0; i < 224; i++ {
			t.Forward(392 - float64(i))
			t.Turn(121)
		}
	}).Draw(m)

	gfx.SavePNG("/tmp/gfx-readme-examples-turtle.png", m)
}
