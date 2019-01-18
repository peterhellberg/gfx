# gfx

[![Build Status](https://travis-ci.org/peterhellberg/gfx.svg?branch=master)](https://travis-ci.org/peterhellberg/gfx)
[![Go Report Card](https://goreportcard.com/badge/github.com/peterhellberg/gfx?style=flat)](https://goreportcard.com/report/github.com/peterhellberg/gfx)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/gfx)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/gfx#license-mit)

Convenience package for dealing with graphics in my pixel drawing experiments.

#### :warning: NO STABILITY GUARANTEES AS OF YET :warning:

## Geometry

The geometry types is based on those found in <https://github.com/faiface/pixel> (but indended for use without Pixel)

### `gfx.Vec`

Vec is a 2D vector type with X and Y coordinates.

### `gfx.Rect`

Rect is a 2D rectangle aligned with the axes of the coordinate system. It is defined by two points, Min and Max.

### `gfx.Matrix`

Matrix is a 2x3 affine matrix that can be used for all kinds of spatial transforms, such as movement, scaling and rotations.

## Line drawing algorithms

### Bresenham's line algorithm

`gfx.DrawBresenhamLine` draws a line using [Bresenham's line algorithm](http://en.wikipedia.org/wiki/Bresenham's_line_algorithm).

## Polygons

A `gfx.Polygon` is represented by a list of vectors.
There is also `gfx.Polyline` which is a slice of polygons forming a line.

## Turtle drawing :turtle:

`gfx.Turtle` is a small Turtle inspired drawing type. (`Resize`, `Turn`, `Move`, `Forward`, `Draw`)

<https://www.cse.wustl.edu/~taoju/research/TurtlesforCADRevised.pdf>

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

### Errors

The `gfx.Error` type is a string that implements the `error` interface.

> If you are using [Ebiten](https://github.com/hajimehoshi/ebiten) then you can return the provided `gfx.ErrDone` error to exit its run loop.

### HTTP

You can use `gfx.GetPNG` to download and decode a PNG given an URL.

### Log

I find that it is fairly common for me to do some logging driven development
when experimenting with graphical effects, so I've included `gfx.Log`,
`gfx.Dump`, `gfx.Printf` and `gfx.Sprintf` in this package.

### Resizing images

You can use `gfx.ResizeImage` to resize an image. (nearest neighbor, mainly useful for pixelated graphics)

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
