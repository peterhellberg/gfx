package gfx

func ExampleDrawBresenhamLine() {
	dst := NewPalettedImage(IR(0, 0, 10, 5), Palette1Bit)

	DrawBresenhamLine(dst, V(1, 1), V(8, 3), ColorWhite)

	for y := 0; y < dst.Bounds().Dy(); y++ {
		for x := 0; x < dst.Bounds().Dx(); x++ {
			switch dst.Index(x, y) {
			case 0:
				Printf("_")
			case 1:
				Printf("X")
			}
		}
		Printf("\n")
	}

	// Output:
	//__________
	//_XX_______
	//___XXXX___
	//_______XX_
	//__________
}
