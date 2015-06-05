// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"image"
	"image/color"
)

type Image struct {
	// Pix holds the image's pixels, as pixel values in native-endian order format. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*PixelSize].
	Pix DataSlice
	// Stride is the Pix stride (in bytes, must align with PixelSize)
	// between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle

	Channels int
	DataType DataType
}

func NewImage(r image.Rectangle, channels int, dataType DataType) *Image {
	return &Image{
		Pix:      make([]byte, r.Dy()*r.Dx()*channels*dataType.ByteSize()),
		Stride:   r.Dx() * channels * dataType.ByteSize(),
		Rect:     r,
		Channels: channels,
		DataType: dataType,
	}
}

func (p *Image) Bounds() image.Rectangle {
	return p.Rect
}

func (p *Image) ColorModel() color.Model {
	return makeModelFunc(p.Channels, p.DataType)
}

func (p *Image) At(x, y int) color.Color {
	panic("TODO")
}

func (p *Image) PixelAt(x, y int) DataSlice {
	if !(image.Point{x, y}.In(p.Rect)) {
		return nil
	}
	i, n := p.PixOffset(x, y), p.PixSize()
	return p.Pix[i:][:n]
}

func (p *Image) Set(x, y int, c color.Color) {
	panic("TODO")
}

func (p *Image) SetPixel(x, y int, c DataSlice) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i, n := p.PixOffset(x, y), p.PixSize()
	copy(p.Pix[i:][:n], c)
}

func (p *Image) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*2
}

func (p *Image) PixSize() int {
	return p.Channels * p.DataType.ByteSize()
}

func (p *Image) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &Image{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &Image{
		Pix:      p.Pix[i:],
		Stride:   p.Stride,
		Rect:     r,
		Channels: p.Channels,
		DataType: p.DataType,
	}
}

func (p *Image) StdImage() image.Image {
	switch {
	case p.Channels == 1 && p.DataType == GDT_Byte:
		return &image.Gray{
			Pix:    p.Pix,
			Stride: p.Stride,
			Rect:   p.Rect,
		}
	case p.Channels == 1 && p.DataType == GDT_UInt16:
		return &image.Gray16{
			Pix:    p.Pix,
			Stride: p.Stride,
			Rect:   p.Rect,
		}
	}
	return p
}
