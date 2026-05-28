package gfx

import (
	"image/color"
	"testing"
)

func TestNewBatch(t *testing.T) {
	b := NewBatch(nil, nil)

	if b.mat != IM {
		t.Fatalf("unexpected matrix")
	}
}

func TestBatchClear(t *testing.T) {
	b := NewBatch(&TrianglesData{}, nil)

	b.Clear()
}

func TestBatchDraw(t *testing.T) {
	b := NewBatch(nil, nil)

	m := NewImage(32, 32)

	b.Draw(NewDrawTarget(m))
}

func TestBatchMakeTriangles(t *testing.T) {
	b := NewBatch(nil, nil)

	b.MakeTriangles(&TrianglesData{})
}

func TestBatchMakePicture(t *testing.T) {
	b := NewBatch(nil, nil)

	b.MakePicture(nil)
}

func TestBatchTrianglesDraw(t *testing.T) {
	bt := &batchTriangles{
		tri: &TrianglesData{},
		tmp: MakeTrianglesData(0),
		dst: NewBatch(&TrianglesData{}, nil),
	}

	bt.Draw()
}

// TestBatchAccumulatesAndProjects exercises the full Batch path: a
// triangle made through Batch.MakeTriangles, with the Batch's matrix
// set to a translation, should land in the container at the translated
// positions while preserving the per-vertex color.
func TestBatchAccumulatesAndProjects(t *testing.T) {
	container := &TrianglesData{}
	b := NewBatch(container, nil)
	b.SetMatrix(IM.Moved(V(10, 20)))

	src := &TrianglesData{
		Vx(V(0, 0), ColorRed),
		Vx(V(1, 0), ColorGreen),
		Vx(V(0, 1), ColorBlue),
	}

	bt := b.MakeTriangles(src)
	bt.Draw()

	if got, want := container.Len(), 3; got != want {
		t.Fatalf("container.Len() = %d, want %d", got, want)
	}

	wantPositions := []Vec{V(10, 20), V(11, 20), V(10, 21)}
	wantColors := []color.NRGBA{ColorRed, ColorGreen, ColorBlue}
	for i, want := range wantPositions {
		if got := container.Position(i); got != want {
			t.Fatalf("container.Position(%d) = %v, want %v", i, got, want)
		}
		if got := container.Color(i); got != wantColors[i] {
			t.Fatalf("container.Color(%d) = %v, want %v", i, got, wantColors[i])
		}
	}
}

// TestBatchClearEmptiesContainer verifies Clear removes everything
// the Batch has accumulated.
func TestBatchClearEmptiesContainer(t *testing.T) {
	container := &TrianglesData{}
	b := NewBatch(container, nil)

	bt := b.MakeTriangles(&TrianglesData{
		Vx(V(0, 0), ColorRed),
		Vx(V(1, 0), ColorRed),
		Vx(V(0, 1), ColorRed),
	})
	bt.Draw()

	if container.Len() == 0 {
		t.Fatal("expected container to have triangles before Clear")
	}

	b.Clear()

	if got := container.Len(); got != 0 {
		t.Fatalf("container.Len() after Clear = %d, want 0", got)
	}
}

// TestBatchSetMatrixAppliesToSubsequentDraws verifies the documented
// behavior: previously-accumulated triangles keep their projection
// when SetMatrix is called, only later draws see the new matrix.
func TestBatchSetMatrixAppliesToSubsequentDraws(t *testing.T) {
	container := &TrianglesData{}
	b := NewBatch(container, nil)

	src := &TrianglesData{
		Vx(V(0, 0), ColorRed),
		Vx(V(1, 0), ColorRed),
		Vx(V(0, 1), ColorRed),
	}

	// First batch with identity.
	b.MakeTriangles(src).Draw()

	// Now move and add the same triangle again.
	b.SetMatrix(IM.Moved(V(100, 100)))
	b.MakeTriangles(src).Draw()

	if got, want := container.Len(), 6; got != want {
		t.Fatalf("container.Len() = %d, want %d", got, want)
	}

	// First 3 vertices were captured under the identity matrix.
	if got := container.Position(0); got != V(0, 0) {
		t.Fatalf("container.Position(0) = %v, want %v (matrix change must not retroact)", got, V(0, 0))
	}
	// Last 3 saw the translation.
	if got := container.Position(3); got != V(100, 100) {
		t.Fatalf("container.Position(3) = %v, want %v (new matrix must apply)", got, V(100, 100))
	}
}

// TestBatchMakePictureWrongPicturePanics pins the documented contract:
// MakePicture must be called with the same Picture the Batch was
// constructed with.
func TestBatchMakePictureWrongPicturePanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected MakePicture(differentPic) to panic")
		}
	}()

	b := NewBatch(&TrianglesData{}, fakePicture{})
	b.MakePicture(fakePicture{tag: "different"})
}

type fakePicture struct{ tag string }

func (fakePicture) Bounds() Rect { return R(0, 0, 1, 1) }
