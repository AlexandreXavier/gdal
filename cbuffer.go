// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

//#include <stdlib.h>
import "C"
import (
	"errors"
	"reflect"
	"unsafe"
)

type CBuffer struct {
	dontResize bool
	cptr       unsafe.Pointer
	data       []byte
}

func NewCBuffer(size int, dontResize ...bool) *CBuffer {
	if size <= 0 {
		panic("gdal: NewCBuffer, bad size")
	}
	p := new(CBuffer)
	p.cptr = C.malloc(C.size_t(size))
	p.data = (*[1 << 30]byte)(p.cptr)[0:size:size]
	if len(dontResize) > 0 {
		p.dontResize = dontResize[0]
	}
	return p
}

func (p *CBuffer) Release() {
	if p != nil {
		if p.cptr != nil {
			C.free(p.cptr)
		}
		p.cptr = nil
		p.data = nil
	}
}

func (p *CBuffer) Resize(size int) error {
	if size <= 0 {
		return errors.New("gdal: CBuffer.Resize, bad size!")
	}
	if p.dontResize {
		return errors.New("gdal: CBuffer.Resize, donot resize!")
	}
	p.Release()
	p.cptr = C.malloc(C.size_t(size))
	p.data = (*[1 << 30]byte)(p.cptr)[0:size:size]
	return nil
}

func (p *CBuffer) Size() int {
	return len(p.data)
}

func (p *CBuffer) Data() []byte {
	return p.data
}

func (p *CBuffer) Own(d []byte) bool {
	if cap(d) == 0 {
		return false
	}
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&d))
	if a, b := hdr.Data, uintptr(p.cptr); a < b || a >= b {
		return false
	}
	if a, b := hdr.Data+uintptr(hdr.Cap), uintptr(p.cptr); a < b || a >= b {
		return false
	}
	return true
}
