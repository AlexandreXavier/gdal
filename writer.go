// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"image"
	"io"
)

type Options struct {
	DriverName string
	Projection string
	Transform  [6]float64
	ExtOptions map[string]string
}

// Encode writes the image m to w in GDAL format.
func Encode(w io.Writer, m image.Image, opt *Options) error {
	panic("TODO")
}
