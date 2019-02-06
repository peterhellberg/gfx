package gfx

func ExampleCmplxSin() {
	Dump(
		CmplxSin(complex(1, 2)),
		CmplxSin(complex(2, 3)),
		CmplxSin(complex(4, 5)),
	)

	// Output:
	// (3.165778513216168+1.9596010414216063i)
	// (9.154499146911428-4.168906959966565i)
	// (-56.16227422023235-48.50245524177091i)
}

func ExampleCmplxSinh() {
	Dump(
		CmplxSinh(complex(1, 2)),
		CmplxSinh(complex(2, 3)),
		CmplxSinh(complex(4, 5)),
	)

	// Output:
	// (-0.4890562590412937+1.4031192506220405i)
	// (-3.59056458998578+0.5309210862485197i)
	// (7.741117553247741-26.18652736460921i)
}

func ExampleCmplxCos() {
	Dump(
		CmplxCos(complex(1, 2)),
		CmplxCos(complex(2, 3)),
		CmplxCos(complex(4, 5)),
	)

	// Output:
	// (2.0327230070196656-3.0518977991518i)
	// (-4.189625690968807-9.109227893755337i)
	// (-48.506859457844584+56.15717492513018i)
}

func ExampleCmplxCosh() {
	Dump(
		CmplxCosh(complex(1, 2)),
		CmplxCosh(complex(2, 3)),
		CmplxCosh(complex(4, 5)),
	)

	// Output:
	// (-0.64214812471552+1.068607421382778i)
	// (-3.7245455049153224+0.5118225699873846i)
	// (7.746313007403075-26.168964053872834i)
}

func ExampleCmplxTan() {
	Dump(
		CmplxTan(complex(1, 2)),
		CmplxTan(complex(2, 3)),
		CmplxTan(complex(4, 5)),
	)

	// Output:
	// (0.033812826079896684+1.0147936161466335i)
	// (-0.0037640256415042484+1.0032386273536098i)
	// (8.983477646971559e-05+1.0000132074347845i)
}

func ExampleCmplxTanh() {
	Dump(
		CmplxTanh(complex(1, 2)),
		CmplxTanh(complex(2, 3)),
		CmplxTanh(complex(4, 5)),
	)

	// Output:
	// (1.16673625724092-0.24345820118572523i)
	// (0.965385879022133-0.009884375038322494i)
	// (1.0005630461157933-0.00036520305451130414i)
}

func ExampleCmplxPow() {
	Dump(
		CmplxPow(complex(1, 2), complex(2, 3)),
		CmplxPow(complex(4, 5), complex(5, 6)),
	)

	// Output:
	// (-0.015132672422722659-0.179867483913335i)
	// (-49.59108992764897+4.323851372977011i)
}

func ExampleCmplxSqrt() {
	Dump(
		CmplxSqrt(complex(1, 2)),
		CmplxSqrt(complex(2, 3)),
		CmplxSqrt(complex(4, 5)),
	)

	// Output:
	// (1.272019649514069+0.7861513777574233i)
	// (1.6741492280355401+0.8959774761298381i)
	// (2.280693341665298+1.096157889501519i)
}

func ExampleCmplxPhase() {
	Dump(
		CmplxPhase(complex(1, 2)),
		CmplxPhase(complex(2, 3)),
		CmplxPhase(complex(4, 5)),
	)

	// Output:
	// 1.1071487177940904
	// 0.982793723247329
	// 0.8960553845713439
}
