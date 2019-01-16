package gfx

import "image"

// LayerData is the data for a layer
type LayerData []int

// Size returns the size of the layer data given the number of columns.
func (ld LayerData) Size(cols int) image.Point {
	l := len(ld)

	if l < cols {
		return Pt(cols, 1)
	}

	rows := l / cols

	if rows*cols == l {
		return Pt(cols, rows)
	}

	if rows%cols > 0 {
		rows++
	}

	return Pt(cols, rows)
}
