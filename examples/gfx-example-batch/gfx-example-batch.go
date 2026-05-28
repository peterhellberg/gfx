package main

import "github.com/peterhellberg/gfx"

func main() {
	var (
		p    = gfx.PaletteInk
		dst  = gfx.NewPaletted(1024, 256, p, p[0])
		b    = gfx.NewBatch(&gfx.TrianglesData{}, nil)
		star = makeStar(12, 5, 5)
	)

	for i := range 120 {
		c := p.Color(1 + (i*3)%(p.Len()-1))

		for j := range *star {
			(*star)[j].Color = c
		}

		b.SetMatrix(gfx.IM.
			RotatedDegrees(gfx.ZV, float64(i)*7).
			Moved(gfx.IV(36+(i%15)*68, 16+(i/15)*32)),
		)

		b.MakeTriangles(star).Draw()
	}

	b.Draw(gfx.NewDrawTarget(dst))

	gfx.SavePNG("gfx-example-batch.png", dst)
}

func makeStar(outer, inner float64, tips int) *gfx.TrianglesData {
	n := tips * 2

	pt := func(i int) gfx.Vec {
		r := outer
		if i%2 == 1 {
			r = inner
		}
		a := -gfx.Pi/2 + float64(i)*gfx.Pi/float64(tips)

		return gfx.V(r*gfx.MathCos(a), r*gfx.MathSin(a))
	}

	var td gfx.TrianglesData

	for i := range n {
		td = append(td,
			gfx.Vx(gfx.V(0, 0)),
			gfx.Vx(pt(i)),
			gfx.Vx(pt((i+1)%n)),
		)
	}

	return &td
}
