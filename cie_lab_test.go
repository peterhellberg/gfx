package gfx

func ExampleLab() {
	var (
		rgba   = ColorRGBA(255, 0, 0, 255)
		xyz    = ColorToXYZ(rgba)
		hunter = xyz.HunterLab(XYZReference2.D65)
		cieLab = xyz.CIELab(XYZReference2.D65)
	)

	Dump(
		"RGBA",
		rgba,
		"XYZ",
		xyz,
		"Hunter",
		hunter,
		"CIE-L*ab",
		cieLab,
	)

	// Output:
	//
	// RGBA
	// {R:255 G:0 B:0 A:255}
	// XYZ
	// {X:41.24 Y:21.26 Z:1.9300000000000002}
	// Hunter
	// {L:46.10856753359401 A:82.7190894239167 B:28.333423774179554}
	// CIE-L*ab
	// {L:53.23288178584246 A:80.10930952982204 B:67.22006831026425}
	//
}
