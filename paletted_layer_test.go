package gfx

import (
	"image"
	"testing"
)

func TestPalettedLayer(t *testing.T) {
	var palette = Palette{
		{0x08, 0x18, 0x20, 0xFF},
		{0x34, 0x68, 0x56, 0xFF},
		{0x88, 0xC0, 0x70, 0xFF},
		{0xE0, 0xF8, 0xD0, 0xFF},
	}

	tileset := &PalettedTileset{
		Size: Pt(4, 4),
		Images: PalettedImages{
			NewPixImage([]uint8{
				0, 1, 2, 3,
				0, 1, 2, 3,
				0, 1, 2, 3,
				0, 1, 2, 3,
			}, 4, palette),
			NewPixImage([]uint8{
				0, 2, 2, 0,
				2, 2, 2, 2,
				2, 2, 2, 2,
				0, 2, 2, 0,
			}, 4, palette),
			NewPixImage([]uint8{
				3, 2, 3, 2,
				2, 3, 2, 2,
				3, 2, 3, 2,
				2, 3, 2, 3,
			}, 4, palette),
		},
	}

	pg := &PalettedLayer{tileset, 3, []int{
		2, 1, 2,
		1, 0, 1,
	}}

	b := pg.Bounds()

	if got, want := b.Dx(), 12; got != want {
		t.Fatalf("b.Dx() = %d, want %d", got, want)
	}

	if got, want := b.Dy(), 8; got != want {
		t.Fatalf("b.Dy() = %d, want %d", got, want)
	}

	var _ tiledPalettedImage = pg
}

type tiledPalettedImage interface {
	TileAt(x, y int) image.PalettedImage
	TileSize() image.Point
	image.PalettedImage
}

func TestPalettedLayerBounds(t *testing.T) {
	l := &PalettedLayer{
		PalettedTileset: &PalettedTileset{
			Size: Pt(10, 5),
		},
		Width: 3,
		Data: LayerData{
			0, 0, 0,
			1,
		}}

	t.Log(l.Bounds())
}
