package gfx

import (
	"testing"
)

func TestDegreesRadians(t *testing.T) {
	for d, want := range map[Degrees]float64{
		1:   0.017453292519943295,
		30:  0.5235987755982988,
		60:  1.0471975511965976,
		90:  1.5707963267948966,
		120: 2.0943951023931953,
	} {
		if got := d.Radians(); got != want {
			t.Fatalf("d.Radians() = %#v, want %#v", got, want)
		}
	}
}
