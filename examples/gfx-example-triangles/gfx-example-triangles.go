package main

import "github.com/peterhellberg/gfx"

var p = gfx.PaletteFamicube

func main() {
	n := 50

	dst := gfx.NewPaletted(320, 448, p, p.Color(n+7))

	t := gfx.NewDrawTarget(dst)

	t.MakeTriangles(&gfx.TrianglesData{
		vx(64, 6, n+1), vx(6, 122, n+2), vx(302, 122, n+3),
		vx(6, 150, n+4), vx(150, 150, n+5), vx(120, 436, n+6),
	}).Draw()

	gfx.SavePNG("/tmp/gfx-triangles.png", dst)
}

func vx(x, y float64, n int) gfx.Vertex {
	return gfx.Vertex{Position: gfx.V(x, y), Color: p.Color(n)}
}
