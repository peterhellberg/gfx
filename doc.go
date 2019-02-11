/*

Package gfx is a convenience package for dealing with graphics in my pixel drawing experiments.

My experiments are often published under https://gist.github.com/peterhellberg


Geometry and Transformation

The geometry and transformation types is based on those found in https://github.com/faiface/pixel
(but indended for use without Pixel)

Vec

gfx.Vec is a 2D vector type with X and Y coordinates.

Rect

gfx.Rect is a 2D rectangle aligned with the axes of the coordinate system. It is defined by two points, Min and Max.

Matrix

gfx.Matrix is a 2x3 affine matrix that can be used for all kinds of spatial transforms, such as movement, scaling and rotations.

		package main

		import "github.com/peterhellberg/gfx"

		var en4 = gfx.PaletteEN4

		func main() {
			a := &gfx.Animation{Delay: 10}

			c := gfx.V(128, 128)

			p := gfx.Polygon{
				{50, 50},
				{50, 206},
				{128, 96},
				{206, 206},
				{206, 50},
			}

			for d := 0.0; d < 360; d += 2 {
				m := gfx.NewPaletted(256, 256, en4, en4.Color(3))

				matrix := gfx.IM.RotatedDegrees(c, d)

				gfx.DrawPolygon(m, p.Project(matrix), 0, en4.Color(2))
				gfx.DrawPolygon(m, p.Project(matrix.Scaled(c, 0.5)), 0, en4.Color(1))

				gfx.DrawCircleFilled(m, c, 5, en4.Color(0))

				a.AddPalettedImage(m)
			}

			a.SaveGIF("/tmp/gfx-readme-examples-matrix.gif")
		}


Line drawing algorithms

Drawing lines is fairly common in my experiments so
I've included Bresenham's line algorithm in this package.

Bresenham's line algorithm

gfx.DrawLineBresenham draws a line using Bresenham's line algorithm.

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


Polygons

A gfx.Polygon is represented by a list of vectors. There is also gfx.Polyline which is a slice of polygons forming a line.

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


Triangles

Triangles can be drawn to an image using a *gfx.DrawTarget.

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

Turtle drawing

gfx.Turtle is a small Turtle inspired drawing type. (Resize, Turn, Move, Forward, Draw)

https://www.cse.wustl.edu/~taoju/research/TurtlesforCADRevised.pdf

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

Animation

There is rudimentary support for making animations using gfx.Animation, the animations can then be encoded into GIF.

		package main

		import "github.com/peterhellberg/gfx"

		func main() {
			a := &gfx.Animation{}
			p := gfx.PaletteEDG36

			var fireflower = []uint8{
				0, 1, 1, 1, 1, 1, 1, 0,
				1, 1, 2, 2, 2, 2, 1, 1,
				1, 2, 3, 3, 3, 3, 2, 1,
				1, 1, 2, 2, 2, 2, 1, 1,
				0, 1, 1, 1, 1, 1, 1, 0,
				0, 0, 0, 4, 4, 0, 0, 0,
				0, 0, 0, 4, 4, 0, 0, 0,
				4, 4, 0, 4, 4, 0, 4, 4,
				0, 4, 0, 4, 4, 0, 4, 0,
				0, 4, 4, 4, 4, 4, 4, 0,
				0, 0, 4, 4, 4, 4, 0, 0,
			}

			for i := 0; i < len(p)-4; i++ {
				t := gfx.NewTile(p[i:i+4], 8, fireflower)

				a.AddPalettedImage(gfx.NewScaledPalettedImage(t, 20))
			}

			a.SaveGIF("/tmp/gfx-readme-examples-animation.gif")
		}


Colors

You can construct new colors using gfx.ColorRGBA and gfx.ColorWithAlpha.

Default colors

There are a few default colors in this package, convenient when you just want to experiment, for more ambitious projects I suggest creating a gfx.Palette (or even use one of the included palettes).

		gfx.ColorBlack
		gfx.ColorWhite
		gfx.ColorTransparent
		gfx.ColorOpaque
		gfx.ColorRed
		gfx.ColorGreen
		gfx.ColorBlue
		gfx.ColorCyan
		gfx.ColorMagenta
		gfx.ColorYellow

Palettes

There are a number of palettes in the gfx package, most of them are found in the Lospec Palette List.

Errors

The gfx.Error type is a string that implements the error interface.

    If you are using Ebiten then you can return the provided gfx.ErrDone error to exit its run loop.

HTTP

You can use gfx.GetPNG to download and decode a PNG given an URL.

Log

I find that it is fairly common for me to do some logging driven development when experimenting with graphical effects, so I've included gfx.Log, gfx.Dump, gfx.Printf and gfx.Sprintf in this package.

Math

I have included a few functions that call functions in the math package.

There is also gfx.Sign, gfx.Clamp and gfx.Lerp functions for float64.

Cmplx

I have included a few functions that call functions in the cmplx package.

Reading files

It is fairly common to read files in my experiments, so I've included gfx.ReadFile and gfx.ReadJSON in this package.

Resizing images

You can use gfx.ResizeImage to resize an image. (nearest neighbor, mainly useful for pixelated graphics)

Noise

Different types of noise is often used in procedural generation.

SimplexNoise

SimplexNoise is a speed-improved simplex noise algorithm for 2D, 3D and 4D.

*/
package gfx
