package main

import "github.com/peterhellberg/gfx"

func main() {
	m := gfx.NewImage(1024, 256)

	gfx.EachPixel(m.Bounds(), func(x, y int) {
		sdg := gfx.SignedDistanceGradient{
			gfx.IV(x, y),
		}

		p := gfx.PaletteNight16
		v := sdg.Circle(450)

		if v.X < 0 {
			p = gfx.PaletteEDG16
		}

		m.Set(x, y, p.At(v.Y*v.Z))
	})

	gfx.SavePNG("gfx-example-sdg.png", m)
}
