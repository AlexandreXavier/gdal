// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"image"
	"io"
)

// Decode reads a GDAL image from r and returns it as an image.Image.
func Decode(r io.Reader) (image.Image, error) {
	panic("TODO")
}

// DecodeConfig returns the color model and dimensions of a GDAL image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (image.Config, error) {
	panic("TODO")
}

func init() {
	image.RegisterFormat("gdal", "????", Decode, DecodeConfig)
}
