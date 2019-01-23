package gfx

import (
	"image"
	"testing"
)

func TestLayerData(t *testing.T) {
	for _, tc := range []struct {
		n    int
		data LayerData
		want image.Point
	}{
		{3, LayerData{}, Pt(3, 1)},
		{3, LayerData{0, 1}, Pt(3, 1)},
		{3, LayerData{0, 1, 2, 3}, Pt(3, 2)},
		{3, LayerData{0, 1, 2, 3, 4, 5, 6}, Pt(3, 3)},
		{4, LayerData{0, 1, 2, 3, 4, 5, 6}, Pt(4, 2)},
		{2, LayerData{0, 1, 2, 3, 4, 5, 6}, Pt(2, 4)},
		{3, LayerData{0, 1, 2, 3, 4, 5}, Pt(3, 2)},
	} {
		if got := tc.data.Size(tc.n); !got.Eq(tc.want) {
			t.Errorf("%v.Size(%d) = %v, want %v", tc.data, tc.n, got, tc.want)
		}
	}
}

func TestDataOffset(t *testing.T) {
	for _, tc := range []struct {
		width int
		size  image.Point
		input image.Point
		want  int
	}{
		{10, Pt(10, 10), Pt(20, 5), 70},
		{30, Pt(30, 5), Pt(20, 10), 320},
	} {
		l := &Layer{Width: tc.width, Tileset: &Tileset{Size: tc.size}}

		if got := l.dataOffset(tc.input.X, tc.input.Y); got != tc.want {
			t.Fatalf("l.indexAt(%d, %d) = %d, want %d",
				tc.input.X, tc.input.Y, got, tc.want)
		}
	}
}
