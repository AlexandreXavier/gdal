// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"image"
	"reflect"
)

// LoadConfig returns the color model and dimensions of a GDAL image without
// decoding the entire image.
func LoadConfig(filename string) (config image.Config, err error) {
	f, err := OpenDataset(filename, GA_ReadOnly)
	if err != nil {
		return
	}
	defer f.Close()

	config.ColorModel = ColorModel(f._Channels, f._DataType)
	config.Width, config.Height = f._Width, f._Height
	return
}

// Load reads a GDAL image from file and returns it as an image.Image.
func Load(filename string, buffer ...[]byte) (m image.Image, err error) {
	f, err := OpenDataset(filename, GA_ReadOnly)
	if err != nil {
		return
	}
	defer f.Close()

	p := NewMemPImage(image.Rect(0, 0, f._Width, f._Height), f._Channels, f._DataType)
	if err = f.ReadToBuf(p.XRect, p.XPix, p.XStride); err != nil {
		return
	}

	if p.XChannels == 1 && p.XDataType == reflect.Uint8 {
		return &image.Gray{
			Pix:    p.XPix,
			Stride: p.XStride,
			Rect:   p.XRect,
		}, nil
	}
	if p.XChannels == 4 && p.XDataType == reflect.Uint8 {
		return &image.RGBA{
			Pix:    p.XPix,
			Stride: p.XStride,
			Rect:   p.XRect,
		}, nil
	}
	if p.XChannels == 1 && p.XDataType == reflect.Uint16 {
		if isLittleEndian {
			p.XPix.SwapEndian(p.XDataType)
		}
		return &image.Gray16{
			Pix:    p.XPix,
			Stride: p.XStride,
			Rect:   p.XRect,
		}, nil
	}
	if p.XChannels == 4 && p.XDataType == reflect.Uint16 {
		if isLittleEndian {
			p.XPix.SwapEndian(p.XDataType)
		}
		return &image.RGBA64{
			Pix:    p.XPix,
			Stride: p.XStride,
			Rect:   p.XRect,
		}, nil
	}

	m = p.StdImage()
	return
}

// LoadImage reads a GDAL image from file and returns it as an Image.
func LoadImage(filename string, buffer ...[]byte) (m *MemPImage, err error) {
	f, err := OpenDataset(filename, GA_ReadOnly)
	if err != nil {
		return
	}
	defer f.Close()

	m = NewMemPImage(image.Rect(0, 0, f._Width, f._Height), f._Channels, f._DataType)
	if err = f.ReadToBuf(m.XRect, m.XPix, m.XStride); err != nil {
		return
	}
	return
}
