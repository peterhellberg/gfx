package gfx

import "image/draw"

// Blocks is a slice of blocks.
type Blocks []Block

// Add appends a Block to the slice.
func (blocks *Blocks) Add(b Block) {
	*blocks = append(*blocks, b)
}

// AddNewBlock creates a new Block and appends it to the slice.
func (blocks *Blocks) AddNewBlock(pos, size Vec3, ic BlockColor) {
	blocks.Add(NewBlock(pos, size, ic))
}

// DrawPolygons draws all of the blocks on the dst image.
// (using the shape, top and left polygons at the given origin)
func (blocks Blocks) DrawPolygons(dst draw.Image, origin Vec3) {
	for _, block := range blocks {
		shape, top, left, _ := block.Polygons(origin)

		shape.Fill(dst, block.Color.Medium)
		left.Fill(dst, block.Color.Dark)
		top.Fill(dst, block.Color.Light)
	}
}
