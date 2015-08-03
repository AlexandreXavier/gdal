// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

//#include <stdlib.h>
import "C"
import (
	"image"
	"os"
	"unsafe"
)

type CBuffer struct {
	cptr unsafe.Pointer
	data []byte
}

func NewCBuffer(size int) *CBuffer {
	p := new(CBuffer)
	p.cptr = C.malloc(C.size_t(size))
	p.data = (*[1 << 30]byte)(p.cptr)[0:size:size]
	return p
}

func (p *CBuffer) Release() error {
	if p != nil {
		if p.cptr != nil {
			C.free(p.cptr)
		}
		p.cptr = nil
		p.data = nil
	}
	return nil
}

func (p *CBuffer) Resize(size int) {
	p.Release()
	p.cptr = C.malloc(C.size_t(size))
	p.data = (*[1 << 30]byte)(p.cptr)[0:size:size]
}

func (p *CBuffer) Size() int {
	return len(p.data)
}

func (p *CBuffer) Data() []byte {
	return p.data
}

// LoadCImage reads a GDAL image from file and returns it as an Image, the m.Pix is in C memory.
func LoadCImage(filename string, cbuf *CBuffer) (m *Image, err error) {
	f, err := OpenDataset(filename, os.O_RDONLY)
	if err != nil {
		return
	}
	defer f.Close()

	if cbuf != nil {
		m = newCImage(cbuf, image.Rect(0, 0, f.Width, f.Height), f.Channels, f.DataType)
		if err = f.ReadToCBuf(m.Rect, m.Pix, m.Stride); err != nil {
			return
		}
	} else {
		m = NewImage(image.Rect(0, 0, f.Width, f.Height), f.Channels, f.DataType)
		if err = f.Read(m.Rect, m.Pix, m.Stride); err != nil {
			return
		}
	}

	return
}
