package gfx

import (
	"cmp"
	"slices"
	"testing"
)

func TestSortSliceInts(t *testing.T) {
	s := []int{3, 1, 4, 1, 5, 9, 2, 6}

	SortSlice(s, cmp.Compare[int])

	if want := []int{1, 1, 2, 3, 4, 5, 6, 9}; !slices.Equal(s, want) {
		t.Fatalf("SortSlice = %v, want %v", s, want)
	}
}

func TestSortSliceCustom(t *testing.T) {
	type item struct {
		Key   string
		Value int
	}
	s := []item{
		{"b", 2},
		{"a", 1},
		{"c", 3},
	}

	SortSlice(s, func(a, b item) int {
		return cmp.Compare(a.Key, b.Key)
	})

	if s[0].Key != "a" || s[1].Key != "b" || s[2].Key != "c" {
		t.Fatalf("SortSlice = %v, want sorted by Key", s)
	}
}
