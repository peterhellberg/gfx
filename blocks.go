package gfx

import (
	"image/draw"
	"sort"
)

// Blocks is a slice of blocks.
type Blocks []Block

// Add appends one or more blocks to the slice of Blocks.
func (blocks *Blocks) Add(bs ...Block) {
	if len(bs) > 0 {
		*blocks = append(*blocks, bs...)
	}
}

// AddNewBlock creates a new Block and appends it to the slice.
func (blocks *Blocks) AddNewBlock(pos, size Vec3, ic BlockColor) {
	blocks.Add(NewBlock(pos, size, ic))
}

// DrawBounds draws the bounds of the blocks on the dst image.
// (using the shape, top and left polygon bounds at the given origin)
func (blocks Blocks) DrawBounds(dst draw.Image, origin Vec3) int {
	var drawCount int

	for _, block := range blocks {
		shape, top, left, right := block.Polygons(origin)

		if shape.Bounds().Overlaps(dst.Bounds()) {
			DrawColor(dst, top.Bounds().Inset(top.Bounds().Dy()/4), block.Color.Light)
			DrawColor(dst, right.Bounds(), block.Color.Medium)
			DrawColor(dst, left.Bounds(), block.Color.Dark)

			drawCount += 3
		}
	}

	return drawCount
}

// DrawPolygons draws all of the blocks on the dst image.
// (using the shape, top and left polygons at the given origin)
func (blocks Blocks) DrawPolygons(dst draw.Image, origin Vec3) int {
	var drawCount int

	for _, block := range blocks {
		shape, top, left, _ := block.Polygons(origin)

		if shape.Bounds().Overlaps(dst.Bounds()) {
			drawCount += shape.Fill(dst, block.Color.Medium)
			drawCount += left.Fill(dst, block.Color.Dark)
			drawCount += top.Fill(dst, block.Color.Light)
		}
	}

	return drawCount
}

// Sort blocks to be drawn starting from max X, max Y and min Z.
func (blocks Blocks) Sort() {
	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].Behind(blocks[j])
	})
}
