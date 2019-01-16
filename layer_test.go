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
