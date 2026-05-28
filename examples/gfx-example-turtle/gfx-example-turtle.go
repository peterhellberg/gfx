package main

import "github.com/peterhellberg/gfx"

func main() {
	p := gfx.PaletteNYX8
	m := gfx.NewImage(1024, 256, p[0])

	x := 74
	n := 21

	ends := []gfx.Vec{}

	for j := range 4 {
		gfx.NewTurtle(gfx.IV(x, 225), func(t *gfx.Turtle) {
			for i := range n {
				t.Width = t.Width + 0.003
				t.Color = p.Color(1 + (i+j*3)%(p.Len()-1))
				t.Forward(196 - float64(i))
				t.Turn(122)

				ends = append(ends, t.Position)
			}
		}).Draw(m)

		x += 250
		n = n * 2
	}

	for _, end := range ends {
		gfx.DrawCircle(m, end, 6, 4, p.Color(3))
	}

	gfx.SavePNG("gfx-example-turtle.png", m)
}
