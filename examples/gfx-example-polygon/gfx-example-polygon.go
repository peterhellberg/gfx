package main

import "github.com/peterhellberg/gfx"

var edg32 = gfx.PaletteEDG32

func main() {
	m := gfx.NewImage(512, 512)

	p := gfx.Polygon{
		{40, 40},
		{240, 60},
		{440, 460},
		{160, 360},
		{180, 140},
	}

	p.Fill(m, edg32.Color(7))

	pc := p.Rect().Center()

	p.EachPixel(m, func(x, y int) {
		pv := gfx.IV(x, y)

		l := pv.To(pc).Len()

		gfx.Mix(m, x, y, edg32.Color(int(l/18)%32))
	})

	for n, v := range p {
		c := edg32.Color(n * 4)

		gfx.DrawCircle(m, v, 15, 8, gfx.ColorWithAlpha(c, 96))
		gfx.DrawCircle(m, v, 16, 1, c)
	}

	gfx.SavePNG("/tmp/gfx-readme-examples-polygon.png", m)
}
