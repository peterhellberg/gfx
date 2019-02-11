package main

import "github.com/peterhellberg/gfx"

var p = gfx.PaletteFamicube

func main() {
	n := 50
	m := gfx.NewPaletted(800, 270, p, p.Color(n+7))
	t := gfx.NewDrawTarget(m)

	t.MakeTriangles(&gfx.TrianglesData{
		vx(64, 16, n+1), vx(6, 142, n+2), vx(302, 142, n+3),
		vx(300, 142, n+4), vx(450, 50, n+5), vx(590, 236, n+6),
		vx(550, 70, n+8), vx(770, 150, n+9), vx(620, 236, n+10),
	}).Draw()

	gfx.SavePNG("gfx-example-triangles.png", m)
}

func vx(x, y float64, n int) gfx.Vertex {
	return gfx.Vertex{Position: gfx.V(x, y), Color: p.Color(n)}
}
