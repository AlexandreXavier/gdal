// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"image"
)

// Encode writes the image m to w in GDAL format.
func Save(filename string, m image.Image, opt *Options) (err error) {
	p := NewImageFrom(m)

	f, err := CreateDataset(filename, p.Rect.Dx(), p.Rect.Dy(), p.Channels, p.DataType, opt)
	if err != nil {
		return
	}
	defer f.Close()

	if err = f.Write(p.Rect, p.Pix, p.Stride); err != nil {
		return
	}
	return
}
