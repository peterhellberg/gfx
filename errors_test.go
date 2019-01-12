package gfx

import "testing"

func TestErrorf(t *testing.T) {

	err := Errorf("foo %d and bar %s", 123, "abc")

	if got, want := err.Error(), "foo 123 and bar abc"; got != want {
		t.Fatalf("err.Error() = %q, want %q", got, want)
	}
}
