package gfx

// Blocks is a slice of blocks.
type Blocks []Block

// Add appends a Block to the slice.
func (bs *Blocks) Add(b Block) {
	*bs = append(*bs, b)
}

// AddNewBlock creates a new Block and appends it to the slice.
func (bs *Blocks) AddNewBlock(pos, size Vec3, ic BlockColor) {
	bs.Add(NewBlock(pos, size, ic))
}

// Block has a position, size and color.
type Block struct {
	Pos   Vec3
	Size  Vec3
	Color BlockColor
}

// NewBlock creates a new Block.
func NewBlock(pos, size Vec3, ic BlockColor) Block {
	return Block{Pos: pos, Size: size, Color: ic}
}

// Box creates a box for the Block.
func (b Block) Box() Box {
	return NewBox(b.Pos, b.Pos.Add(b.Size))
}

// Polygons returns the shape, top, left and right polygons with coordinates based on origin.
func (b Block) Polygons(origin Vec3) (shape, top, left, right Polygon) {
	v := b.screenVecs(origin)

	return v.Shape(), v.Top(), v.Left(), v.Right()
}

func (b Block) screenVecs(origin Vec3) screenVecs {
	sv := newSpaceVec3s(b.Pos, b.Size)

	return screenVecs{
		LeftUp:    spaceToScreen(sv.LeftUp, origin),
		LeftDown:  spaceToScreen(sv.LeftDown, origin),
		FrontDown: spaceToScreen(sv.FrontDown, origin),
		RightDown: spaceToScreen(sv.RightDown, origin),
		RightUp:   spaceToScreen(sv.RightUp, origin),
		BackUp:    spaceToScreen(sv.BackUp, origin),
		BackDown:  spaceToScreen(sv.BackDown, origin),
		FrontUp:   spaceToScreen(sv.FrontUp, origin),
	}
}

type spaceVec3s struct {
	LeftUp    Vec3
	LeftDown  Vec3
	FrontDown Vec3
	RightDown Vec3
	RightUp   Vec3
	BackUp    Vec3
	BackDown  Vec3
	FrontUp   Vec3
}

func newSpaceVec3s(p, s Vec3) spaceVec3s {
	return spaceVec3s{
		LeftUp:    V3(p.X, p.Y+s.Y, p.Z+s.Z),
		LeftDown:  V3(p.X, p.Y+s.Y, p.Z),
		FrontDown: V3(p.X, p.Y, p.Z),
		RightDown: V3(p.X+s.X, p.Y, p.Z),
		RightUp:   V3(p.X+s.X, p.Y, p.Z+s.Z),
		BackUp:    V3(p.X+s.X, p.Y+s.Y, p.Z+s.Z),
		BackDown:  V3(p.X+s.X, p.Y+s.Y, p.Z),
		FrontUp:   V3(p.X, p.Y, p.Z+s.Z),
	}
}

type screenVecs struct {
	LeftUp    Vec
	LeftDown  Vec
	FrontDown Vec
	RightDown Vec
	RightUp   Vec
	BackUp    Vec
	BackDown  Vec
	FrontUp   Vec
}

func (sv screenVecs) Shape() Polygon {
	return Polygon{sv.LeftUp, sv.LeftDown, sv.FrontDown, sv.RightDown, sv.RightUp, sv.BackUp}
}

func (sv screenVecs) Top() Polygon {
	return Polygon{sv.LeftUp, sv.FrontUp, sv.RightUp, sv.BackUp}
}

func (sv screenVecs) Left() Polygon {
	return Polygon{sv.LeftUp, sv.LeftDown, sv.FrontDown, sv.FrontUp}
}

func (sv screenVecs) Right() Polygon {
	return Polygon{sv.FrontUp, sv.FrontDown, sv.RightDown, sv.RightUp}
}

// spaceToScreen converts a 3D space position and origin into a screen position.
func spaceToScreen(space, origin Vec3) Vec {
	h, v := spaceToIso(space)

	// Convert the given 2D isometric coordinates to 2D screen coordinates.
	x := h*origin.Z + origin.X
	y := -(v*origin.Z + origin.Y)

	return V(x, y)
}

// Convert 3D space coordinates to flattened 2D isometric coordinates.
// x and y coordinates are oblique axes separated by 120 degrees.
// h,v are the horizontal and vertical distances from the origin.
func spaceToIso(space Vec3) (h, v float64) {
	x, y := space.X+space.Z, space.Y+space.Z

	h = (x - y) * MathSqrt(3) / 2
	v = (x + y) / 2

	return h, v
}
