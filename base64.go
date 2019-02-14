package gfx

import (
	"bytes"
	"encoding/base64"
	"image"
)

// Base64EncodedPNG encodes the given image into
// a string using base64.StdEncoding.
func Base64EncodedPNG(src image.Image) string {
	var buf bytes.Buffer

	EncodePNG(&buf, src)

	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
