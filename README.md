# gfx

[![Build Status](https://travis-ci.org/peterhellberg/gfx.svg?branch=master)](https://travis-ci.org/peterhellberg/gfx)
[![Go Report Card](https://goreportcard.com/badge/github.com/peterhellberg/gfx?style=flat)](https://goreportcard.com/report/github.com/peterhellberg/gfx)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/gfx)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/gfx#license-mit)

** :warning: NO STABILITY GUARANTEES AS OF YET **

Convenience package for dealing with graphics in my pixel drawing experiments.

## Geometry

The geometry types is based on those found in <https://github.com/faiface/pixel> (but indended for use without Pixel)

### `Vec` a 2D vector type

### `Rect` is a 2D rectangle

## Line algorithms

### Bresenham's line algorithm

<http://en.wikipedia.org/wiki/Bresenham's_line_algorithm>

## `Turtle` drawing :turtle:

A small Turtle inspired drawing type. (`Resize`, `Turn`, `Move`, `Forward`, `Draw`)

<https://www.cse.wustl.edu/~taoju/research/TurtlesforCADRevised.pdf>

## License (MIT)

Copyright (c) 2019 [Peter Hellberg](https://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
