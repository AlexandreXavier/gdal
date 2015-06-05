// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"image/color"
)

type Pixel struct {
	Channels int
	DataType DataType
	Pix      DataView
}

func (c Pixel) RGBA() (r, g, b, a uint32) {
	return
}

func ColorModel(channels int, dataType DataType) color.Model {
	return color.ModelFunc(func(c color.Color) color.Color {
		_ = channels
		_ = dataType
		return c
	})
}

func colorRgbToGray(r, g, b uint32) uint32 {
	y := (299*r + 587*g + 114*b + 500) / 1000
	return y
}

func colorRgbToGrayI32(r, g, b int32) int32 {
	y := (299*r + 587*g + 114*b + 500) / 1000
	return y
}

func colorRgbToGrayF32(r, g, b float32) float32 {
	y := (299*r + 587*g + 114*b + 500) / 1000
	return y
}

func colorRgbToGrayI64(r, g, b int64) int64 {
	y := (299*r + 587*g + 114*b + 500) / 1000
	return y
}

func colorRgbToGrayF64(r, g, b float64) float64 {
	y := (299*r + 587*g + 114*b + 500) / 1000
	return y
}
