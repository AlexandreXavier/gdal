// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"image"
	"os"
	"reflect"
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
func Load(filename string, cbuf ...*CBuffer) (m image.Image, err error) {
	f, err := OpenDataset(filename, os.O_RDONLY)
	if err != nil {
		return
	}
	defer f.Close()

	var p *MemPImage = nil
	if len(cbuf) > 0 && cbuf[0] != nil {
		p = newCImage(cbuf[0], image.Rect(0, 0, f.Width, f.Height), f.Channels, f.DataType)
		if err = f.ReadToCBuf(p.XRect, p.XPix, p.XStride); err != nil {
			return
		}
	} else {
		p = NewMemPImage(image.Rect(0, 0, f.Width, f.Height), f.Channels, f.DataType)
		if err = f.Read(p.XRect, p.XPix, p.XStride); err != nil {
			return
		}
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
func LoadImage(filename string, cbuf ...*CBuffer) (m *MemPImage, err error) {
	f, err := OpenDataset(filename, os.O_RDONLY)
	if err != nil {
		return
	}
	defer f.Close()

	if len(cbuf) > 0 && cbuf[0] != nil {
		m = newCImage(cbuf[0], image.Rect(0, 0, f.Width, f.Height), f.Channels, f.DataType)
		if err = f.ReadToCBuf(m.XRect, m.XPix, m.XStride); err != nil {
			return
		}
	} else {
		m = NewMemPImage(image.Rect(0, 0, f.Width, f.Height), f.Channels, f.DataType)
		if err = f.Read(m.XRect, m.XPix, m.XStride); err != nil {
			return
		}
	}

	return
}

func newCImage(cbuf *CBuffer, r image.Rectangle, channels int, dataType reflect.Kind) *MemPImage {
	p := &MemPImage{
		XMemPMagic: MemPMagic,
		XRect:      r,
		XStride:    r.Dx() * channels * SizeofKind(dataType),
		XChannels:  channels,
		XDataType:  dataType,
	}
	if n := r.Dy() * p.XStride; n > cbuf.Size() {
		if err := cbuf.Resize(n); err != nil {
			panic(err)
		}
	}
	p.XPix = cbuf.Data()
	return p
}
