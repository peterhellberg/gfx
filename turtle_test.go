package gfx

import (
	"testing"
)

func TestTurtleNoOps(t *testing.T) {
	tr := NewTurtle(V(0, 0))

	if got, want := tr.Bounds(), IR(0, 0, 0, 0); got != want {
		t.Fatalf("Bounds() = %v, want %v (empty when no ops)", got, want)
	}
}

func TestTurtleForwardPosition(t *testing.T) {
	tr := NewTurtle(V(10, 10), func(tr *Turtle) {
		tr.Forward(20)
	})

	// Default direction is (0, -1), so 20 steps move Y from 10 to -10.
	if got, want := tr.Position, V(10, -10); got != want {
		t.Fatalf("Position after Forward(20) = %v, want %v", got, want)
	}
}

func TestTurtleTurnAndForward(t *testing.T) {
	tr := NewTurtle(V(0, 0), func(tr *Turtle) {
		tr.Turn(90) // direction rotates from (0, -1) toward (1, 0)
		tr.Forward(10)
	})

	// After turning 90° and moving 10 steps the position should be near (10, 0).
	if got := tr.Position; int(got.X+0.5) != 10 || int(got.Y+0.5) != 0 {
		t.Fatalf("Position after Turn(90) + Forward(10) = %v, want approximately (10, 0)", got)
	}
}
