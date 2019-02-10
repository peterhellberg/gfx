package main

import "github.com/peterhellberg/gfx"

func main() {
	var (
		p      = append(gfx.PaletteGo, gfx.ColorTransparent)
		dst    = gfx.NewPaletted(898, 430, p, gfx.ColorTransparent)
		rect   = gfx.BoundsToRect(dst.Bounds())
		origin = rect.Center().ScaledXY(gfx.V(1.2, -2.5)).Vec3(0.7)
		blocks gfx.Blocks
	)

	for i, bc := range gfx.BlockColorsGo {
		var (
			f    = float64(i) + 0.5
			v    = f * 15
			pos  = gfx.V3(290+(v*3), 8.5*v, -8*(f+2))
			size = gfx.V3(90, 90, 90)
		)

		blocks.AddNewBlock(pos, size, bc)
	}

	blocks.DrawPolygons(dst, origin)

	gfx.SavePNG("gfx-example-blocks.png", dst)
}
