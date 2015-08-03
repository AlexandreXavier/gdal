// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"image"
)

// Encode writes the image m to w in GDAL format.
func Save(filename string, m image.Image, opt *Options) (err error) {
	p, ok := AsMemPImage(m)
	if !ok {
		p = NewMemPImageFrom(m)
	}

	f, err := CreateDataset(filename, p.XRect.Dx(), p.XRect.Dy(), p.XChannels, p.XDataType, opt)
	if err != nil {
		return
	}
	defer f.Close()

	if err = f.Write(p.XRect, p.XPix, p.XStride); err != nil {
		return
	}
	return
}
