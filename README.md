# gfx

[![Build Status](https://travis-ci.org/peterhellberg/gfx.svg?branch=master)](https://travis-ci.org/peterhellberg/gfx)
[![Go Report Card](https://goreportcard.com/badge/github.com/peterhellberg/gfx?style=flat)](https://goreportcard.com/report/github.com/peterhellberg/gfx)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/gfx)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/gfx#license-mit)

Convenience package for dealing with graphics in my pixel drawing experiments.

#### :warning: NO STABILITY GUARANTEES AS OF YET :warning:

## Geometry and Transformation

The geometry and transformation types is based on those found in <https://github.com/faiface/pixel> (but indended for use without Pixel)

### Vec

`gfx.Vec` is a 2D vector type with X and Y coordinates.

### Rect

`gfx.Rect` is a 2D rectangle aligned with the axes of the coordinate system. It is defined by two points, Min and Max.

### Matrix

`gfx.Matrix` is a 2x3 affine matrix that can be used for all kinds of spatial transforms, such as movement, scaling and rotations.

[embedmd]:# (examples/gfx-example-matrix/gfx-example-matrix.go)
```go
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
```

![gfx-readme-examples-matrix](https://user-images.githubusercontent.com/565124/51478881-f8e69a00-1d8c-11e9-92c5-270c767dfc06.gif)

## Line drawing algorithms

### Bresenham's line algorithm

`gfx.DrawLineBresenham` draws a line using [Bresenham's line algorithm](http://en.wikipedia.org/wiki/Bresenham's_line_algorithm).

[embedmd]:# (examples/gfx-example-bresenham-line/gfx-example-bresenham-line.go)
```go
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
```

![gfx-readme-examples-bresenham-line](https://user-images.githubusercontent.com/565124/51472593-3a217e80-1d7a-11e9-902e-6875d3ee6cb8.png)

## Polygons

A `gfx.Polygon` is represented by a list of vectors.
There is also `gfx.Polyline` which is a slice of polygons forming a line.

[embedmd]:# (examples/gfx-example-polygon/gfx-example-polygon.go)
```go
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
```

![gfx-readme-examples-polygon](https://user-images.githubusercontent.com/565124/51088235-61b28e80-175d-11e9-924d-835487277f4a.png)

## Triangles

Triangles can be drawn to an image using a `*gfx.DrawTarget`.

```go
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
```

![gfx-readme-examples-triangles](https://user-images.githubusercontent.com/565124/51729680-fe85fd80-2074-11e9-9079-05b3ef415441.png)

## :turtle: drawing

`gfx.Turtle` is a small Turtle inspired drawing type. (`Resize`, `Turn`, `Move`, `Forward`, `Draw`)

<https://www.cse.wustl.edu/~taoju/research/TurtlesforCADRevised.pdf>

[embedmd]:# (examples/gfx-example-turtle/gfx-example-turtle.go)
```go
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
```

![gfx-readme-examples-turtle](https://user-images.githubusercontent.com/565124/51402174-0ad9fa00-1b4d-11e9-95b9-f5617979f34e.png)

## Animation

There is rudimentary support for making animations using `gfx.Animation`, the animations can then be encoded into GIF.

```go
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
```

![gfx-readme-examples-animation](https://user-images.githubusercontent.com/565124/51402952-437ad300-1b4f-11e9-89f3-292f69f38429.gif)

## Colors

You can construct new colors using `gfx.ColorNRGBA` and `gfx.ColorWithAlpha`.

### Default colors

There are a few default colors in this package, convenient when you just want to experiment,
for more ambitious projects I suggest creating a `gfx.Palette` (or even use one of the included palettes).


| Variable               | Color
|------------------------|---------------------------------------------------------
| `gfx.ColorBlack`       | ![gfx.ColorBlack](https://fakeimg.pl/128x32/000000?text=+)
| `gfx.ColorWhite`       | ![gfx.ColorWhite](https://fakeimg.pl/128x32/FFFFFF?text=+)
| `gfx.ColorTransparent` | ![gfx.ColorTransparent](https://fakeimg.pl/128x32/000000,0/?text=+)
| `gfx.ColorOpaque`      | ![gfx.ColorOpaque](https://fakeimg.pl/128x32/FFFFFF?text=+)
| `gfx.ColorRed`         | ![gfx.ColorRed](https://fakeimg.pl/128x32/FF0000?text=+)
| `gfx.ColorGreen`       | ![gfx.ColorGreen](https://fakeimg.pl/128x32/00FF00?text=+)
| `gfx.ColorBlue`        | ![gfx.ColorBlue](https://fakeimg.pl/128x32/0000FF?text=+)
| `gfx.ColorCyan`        | ![gfx.ColorCyan](https://fakeimg.pl/128x32/00FFFF?text=+)
| `gfx.ColorMagenta`     | ![gfx.ColorMagenta](https://fakeimg.pl/128x32/FF00FF?text=+)
| `gfx.ColorYellow`      | ![gfx.ColorYellow](https://fakeimg.pl/128x32/FFFF00?text=+)

### Palettes

There are a number of palettes in the `gfx` package,
most of them are found in the [Lospec Palette List](https://lospec.com/palette-list/).

| Variable                   | Colors | Lospec Palette
|----------------------------|-------:| -----------------------------------------------------
| `gfx.Palette1Bit`          |      2 |
| `gfx.Palette2BitGrayScale` |      4 | ![2-bit-grayscale](https://lospec.com/palette-list/2-bit-grayscale-32x.png)
| `gfx.PaletteEN4`           |      4 | ![en4](https://lospec.com/palette-list/en4-32x.png)
| `gfx.PaletteARQ4`          |      4 | ![arq4](https://lospec.com/palette-list/arq4-32x.png)
| `gfx.PaletteInk`           |      5 | ![ink](https://lospec.com/palette-list/ink-32x.png)
| `gfx.Palette3Bit`          |      8 | ![3-bit-rgb](https://lospec.com/palette-list/3-bit-rgb-32x.png)
| `gfx.PaletteEDG8`          |      8 | ![endesega-8](https://lospec.com/palette-list/endesga-8-32x.png)
| `gfx.PaletteAmmo8`         |      8 | ![ammo-8](https://lospec.com/palette-list/ammo-8-32x.png)
| `gfx.PaletteNYX8`          |      8 | ![nyx8](https://lospec.com/palette-list/nyx8-32x.png)
| `gfx.Palette15PDX`         |     15 | ![15p-dx](https://lospec.com/palette-list/15p-dx-32x.png)
| `gfx.PaletteCGA`           |     16 | ![color-graphics-adapter](https://lospec.com/palette-list/color-graphics-adapter-32x.png)
| `gfx.PalettePICO8`         |     16 | ![pico-8](https://lospec.com/palette-list/pico-8-32x.png)
| `gfx.PaletteNight16`       |     16 | ![night-16](https://lospec.com/palette-list/night-16-32x.png)
| `gfx.PaletteAAP16`         |     16 | ![aap-16](https://lospec.com/palette-list/aap-16-32x.png)
| `gfx.PaletteArne16`        |     16 | ![arne-16](https://lospec.com/palette-list/arne-16-32x.png)
| `gfx.PaletteEDG16`         |     16 | ![endesega-16](https://lospec.com/palette-list/endesga-16-32x.png)
| `gfx.Palette20PDX`         |     20 | ![20p-dx](https://lospec.com/palette-list/20p-dx-32x.png)
| `gfx.PaletteEDG32`         |     32 | ![endesega-32](https://lospec.com/palette-list/endesga-32-32x.png)
| `gfx.PaletteEDG36`         |     36 | ![endesega-36](https://lospec.com/palette-list/endesga-36-32x.png)
| `gfx.PaletteEDG64`         |     64 | ![endesega-64](https://lospec.com/palette-list/endesga-64-32x.png)
| `gfx.PaletteAAP64`         |     64 | ![aap-64](https://lospec.com/palette-list/aap-64-32x.png)
| `gfx.PaletteFamicube`      |     64 | ![famicube](https://lospec.com/palette-list/famicube-32x.png)
| `gfx.PaletteSplendor128`   |    128 | ![aap-splendor128](https://lospec.com/palette-list/aap-splendor128-32x.png)

## Errors

The `gfx.Error` type is a string that implements the `error` interface.

> If you are using [Ebiten](https://github.com/hajimehoshi/ebiten) then you can return the provided `gfx.ErrDone` error to exit its run loop.

## HTTP

You can use `gfx.GetPNG` to download and decode a PNG given an URL.

## Log

I find that it is fairly common for me to do some logging driven development
when experimenting with graphical effects, so I've included `gfx.Log`,
`gfx.Dump`, `gfx.Printf` and `gfx.Sprintf` in this package.

## Math

I have included a few functions that call functions in the `math` package.

There is also `gfx.Sign`, `gfx.Clamp` and `gfx.Lerp` functions for `float64`.

## Cmplx

I have included a few functions that call functions in the `cmplx` package.

## Reading files

It is fairly common to read files in my experiments, so I've included `gfx.ReadFile` and `gfx.ReadJSON` in this package.

## Resizing images

You can use `gfx.ResizeImage` to resize an image. (nearest neighbor, mainly useful for pixelated graphics)

## Noise

Different types of noise is often used in procedural generation.

### SimplexNoise

SimplexNoise is a speed-improved simplex noise algorithm for 2D, 3D and 4D.

## License (MIT)

Copyright (c) 2019 [Peter Hellberg](https://c7.se)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:
>
> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

<img src="https://data.gopher.se/gopher/viking-gopher.svg" align="right" width="230" height="230">

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
