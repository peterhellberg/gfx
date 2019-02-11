package main

import "github.com/peterhellberg/gfx"

func main() {
	m := gfx.NewImage(1024, 256)
	x := 74
	n := 21

	for j := 0; j < 4; j++ {
		gfx.NewTurtle(gfx.IV(x, 225), func(t *gfx.Turtle) {
			for i := 0; i < n; i++ {
				t.Forward(196 - float64(i))
				t.Turn(122)
			}
		}).Draw(m)

		x += 250
		n = n * 2
	}

	gfx.SavePNG("gfx-example-turtle.png", m)
}
