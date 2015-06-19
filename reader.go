// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"image"
	"os"
)

// LoadConfig returns the color model and dimensions of a GDAL image without
// decoding the entire image.
func LoadConfig(filename string) (config image.Config, err error) {
	f, err := OpenDataset(filename, os.O_RDONLY)
	if err != nil {
		return
	}
	defer f.Close()

	config.ColorModel = ColorModel(f.Channels, f.DataType)
	config.Width, config.Height = f.Width, f.Height
	return
}

// Load reads a GDAL image from file and returns it as an image.Image.
func Load(filename string) (m image.Image, err error) {
	f, err := OpenDataset(filename, os.O_RDONLY)
	if err != nil {
		return
	}
	defer f.Close()

	p := NewImage(image.Rect(0, 0, f.Width, f.Height), f.Channels, f.DataType)
	if err = f.Read(p.Rect, p.Pix, p.Stride); err != nil {
		return
	}

	m = p.StdImage()
	return
}

// LoadImage reads a GDAL image from file and returns it as an Image.
func LoadImage(filename string) (m *Image, err error) {
	f, err := OpenDataset(filename, os.O_RDONLY)
	if err != nil {
		return
	}
	defer f.Close()

	m = NewImage(image.Rect(0, 0, f.Width, f.Height), f.Channels, f.DataType)
	if err = f.Read(m.Rect, m.Pix, m.Stride); err != nil {
		return
	}

	return
}
