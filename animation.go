package gfx

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"io"
	"os"
)

// DefaultAnimationDelay is the default animation delay, in 100ths of a second.
var DefaultAnimationDelay = 50

// Animation represents multiple images.
type Animation struct {
	Frames []*PalettedImage // The successive images.
	Delay  int              // Delay between each of the frames.

	// LoopCount controls the number of times an animation will be
	// restarted during display.
	// A LoopCount of 0 means to loop forever.
	// A LoopCount of -1 means to show each frame only once.
	// Otherwise, the animation is looped LoopCount+1 times.
	LoopCount int
}

// Add a frame to the animation.
func (a *Animation) Add(f *PalettedImage) {
	a.Frames = append(a.Frames, f)
}

// SaveGIF saves the animation to a GIF using the provided file name.
func (a *Animation) SaveGIF(fn string) error {
	w, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer w.Close()

	return a.EncodeGIF(w)
}

// EncodeGIF writes the animation to w in GIF format with the
// given loop count and delay between frames.
func (a *Animation) EncodeGIF(w io.Writer) error {
	var frames []*image.Paletted
	var delays []int

	if a.Delay < 1 {
		a.Delay = DefaultAnimationDelay
	}

	for _, src := range a.Frames {
		dst := image.NewPaletted(src.Bounds(), asColorPalette(src.Palette))

		draw.Draw(dst, dst.Bounds(), src, image.ZP, draw.Src)

		frames = append(frames, dst)
		delays = append(delays, a.Delay)
	}

	return gif.EncodeAll(w, &gif.GIF{
		Image:     frames,
		Delay:     delays,
		LoopCount: a.LoopCount,
	})
}

func asColorPalette(p Palette) color.Palette {
	var cp = make(color.Palette, len(p))

	for i, c := range p {
		cp[i] = c
	}

	return cp
}
