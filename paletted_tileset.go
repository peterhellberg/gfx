package gfx

import "image"

// PalettedTileset is a paletted tileset.
type PalettedTileset struct {
	Size   image.Point    // Size is the size of each tile.
	Images PalettedImages // Images containst all of the images in the tileset.
}

// TilesetData is the raw data in a tileset
type TilesetData [][]uint8

// NewPalettedTileset creates a new paletted tileset.
func NewPalettedTileset(p Palette, s image.Point, td TilesetData) *PalettedTileset {
	ts := &PalettedTileset{Size: s}

	for i := 0; i < len(td); i++ {
		ts.Images = append(ts.Images, NewPixImage(td[i], s.X, p))
	}

	return ts
}
